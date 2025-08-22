package app

import (
	"fmt"
	"images-service/internal/config"
	"images-service/internal/config/env"
	"images-service/internal/infra/grpc_server/controllers/images_controller"
	"images-service/internal/pkg/cloudflare"
	"log"
	"net"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(port int) {
	// Initialize Cloudflare configuration
	cloudflareConfig := &config.CloudflareConfig{
		AccountID:   env.CLOUDFLARE_ACCOUNT_ID,
		APIToken:    env.CLOUDFLARE_API_TOKEN,
		DeliveryURL: env.CLOUDFLARE_DELIVERY_URL,
	}

	// Initialize Cloudflare Images client
	_ = cloudflare.NewImagesClient(cloudflareConfig)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)
	grpcServer := grpc.NewServer(opts...)

	RegisterControllers(grpcServer, cloudflareConfig)
	reflection.Register(grpcServer)

	fmt.Printf("Server is running on port:%d\n", port)
	grpcServer.Serve(lis)
}

func RegisterControllers(grpcServer *grpc.Server, cloudflareConfig *config.CloudflareConfig) {
	// Register only Images controller since we're using only Cloudflare
	images_proto.RegisterImagesServiceServer(grpcServer, images_controller.New(nil, cloudflareConfig))
}
