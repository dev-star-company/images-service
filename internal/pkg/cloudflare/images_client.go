package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"images-service/internal/config"
	"io"
	"mime/multipart"
	"net/http"
)

type ImagesClient struct {
	config     *config.CloudflareConfig
	httpClient *http.Client
}

type UploadResponse struct {
	Success bool `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   struct {
		ID       string `json:"id"`
		Filename string `json:"filename"`
		Uploaded string `json:"uploaded"`
		Meta     struct {
			Key string `json:"key"`
		} `json:"meta"`
		Variants []string `json:"variants"`
	} `json:"result"`
}

type DeleteResponse struct {
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   interface{}   `json:"result"`
}

func NewImagesClient(cfg *config.CloudflareConfig) *ImagesClient {
	return &ImagesClient{
		config:     cfg,
		httpClient: &http.Client{},
	}
}

func (c *ImagesClient) UploadImage(ctx context.Context, imageData []byte, filename string, metadata map[string]string) (*UploadResponse, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Add image file
	fw, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err = fw.Write(imageData); err != nil {
		return nil, err
	}

	// Add metadata
	for key, value := range metadata {
		if err := writer.WriteField(fmt.Sprintf("metadata[%s]", key), value); err != nil {
			return nil, err
		}
	}

	writer.Close()

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/images/v1", c.config.AccountID)
	req, err := http.NewRequestWithContext(ctx, "POST", url, &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var uploadResp UploadResponse
	if err := json.Unmarshal(body, &uploadResp); err != nil {
		return nil, err
	}

	if !uploadResp.Success {
		return &uploadResp, fmt.Errorf("upload failed: %v", uploadResp.Errors)
	}

	return &uploadResp, nil
}

func (c *ImagesClient) DeleteImage(ctx context.Context, imageID string) error {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/images/v1/%s", c.config.AccountID, imageID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete image: status %d", resp.StatusCode)
	}

	return nil
}

func (c *ImagesClient) GetImageURL(imageID, variant string) string {
	if variant == "" {
		variant = "public"
	}
	return fmt.Sprintf("%s/%s/%s", c.config.DeliveryURL, imageID, variant)
}
