package images_controller

import (
	"context"
	"images-service/internal/app/ent"

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
	images_proto.UnimplementedImagesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
