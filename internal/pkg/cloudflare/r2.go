package cloudflare

// PARA APLICAÇÃO FUTURA

// import (
// 	"context"
// 	"io"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/credentials"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// )

// type R2Client struct {
// 	client *s3.Client
// 	bucket string
// }

// func NewR2Client(accountID, accessKey, secretKey, bucket string) *R2Client {
// 	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
// 		return aws.Endpoint{
// 			URL: "https://" + accountID + ".r2.cloudflarestorage.com",
// 		}, nil
// 	})

// 	cfg, _ := config.LoadDefaultConfig(context.TODO(),
// 		config.WithEndpointResolverWithOptions(r2Resolver),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
// 		config.WithRegion("auto"),
// 	)

// 	client := s3.NewFromConfig(cfg)
// 	return &R2Client{client: client, bucket: bucket}
// }

// func (r *R2Client) UploadImage(ctx context.Context, key string, body io.Reader, contentType string) error {
// 	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
// 		Bucket:      aws.String(r.bucket),
// 		Key:         aws.String(key),
// 		Body:        body,
// 		ContentType: aws.String(contentType),
// 	})
// 	return err
// }
