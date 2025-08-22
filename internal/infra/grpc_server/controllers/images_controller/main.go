package images_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/config"
	"images-service/internal/pkg/cloudflare"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
)

type Controller interface {
	images_proto.ImagesServiceServer

	Create(ctx context.Context, in *images_proto.CreateRequest) (*images_proto.CreateResponse, error)
	Get(ctx context.Context, in *images_proto.GetRequest) (*images_proto.GetResponse, error)
	Update(ctx context.Context, in *images_proto.UpdateRequest) (*images_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *images_proto.DeleteRequest) (*images_proto.DeleteResponse, error)
	List(ctx context.Context, in *images_proto.ListRequest) (*images_proto.ListResponse, error)
}

type controller struct {
	Db               *ent.Client
	CloudflareClient *cloudflare.ImagesClient
	images_proto.UnimplementedImagesServiceServer
}

func New(db *ent.Client, cfg *config.CloudflareConfig) *controller {
	return &controller{
		Db:               db,
		CloudflareClient: cloudflare.NewImagesClient(cfg),
	}
}
