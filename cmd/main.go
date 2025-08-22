package main

import (
	"context"
	"fmt"
	"images-service/internal/config"
	"images-service/internal/config/env"
	"images-service/internal/pkg/cloudflare"
)

// ImageService provides methods to interact with Cloudflare Images
type ImageService struct {
	client *cloudflare.ImagesClient
}

func main() {
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}

	fmt.Println("📸 Cloudflare Images Service Ready!")
	fmt.Println("🔗 Import this package in your other services:")
	fmt.Println("   - product_service")
	fmt.Println("   - permission_service")
	fmt.Println("   - Any other Go service")
	fmt.Println("")
	fmt.Println("Usage example:")
	fmt.Println(`   import "images-service/cmd"`)
	fmt.Println(`   service := main.NewImageService()`)
	fmt.Println(`   resp, err := service.Upload(ctx, imageData, "image.jpg", nil)`)
}

// NewImageService creates a new image service instance
func NewImageService() *ImageService {
	// Validate environment on creation
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}

	cloudflareConfig := &config.CloudflareConfig{
		AccountID:   env.CLOUDFLARE_ACCOUNT_ID,
		APIToken:    env.CLOUDFLARE_API_TOKEN,
		DeliveryURL: env.CLOUDFLARE_DELIVERY_URL,
	}

	return &ImageService{
		client: cloudflare.NewImagesClient(cloudflareConfig),
	}
}

// Upload uploads an image to Cloudflare Images
func (s *ImageService) Upload(ctx context.Context, imageData []byte, filename string, metadata map[string]string) (*cloudflare.UploadResponse, error) {
	if metadata == nil {
		metadata = make(map[string]string)
	}
	return s.client.UploadImage(ctx, imageData, filename, metadata)
}

// Delete removes an image from Cloudflare Images
func (s *ImageService) Delete(ctx context.Context, imageID string) error {
	return s.client.DeleteImage(ctx, imageID)
}

// GetURL generates a public URL for an image with optional variant
func (s *ImageService) GetURL(imageID, variant string) string {
	if variant == "" {
		variant = "public"
	}
	return s.client.GetImageURL(imageID, variant)
}

// GetVariantURL generates URLs for different image variants
func (s *ImageService) GetVariantURL(imageID string, variant string) string {
	return s.client.GetImageURL(imageID, variant)
}
