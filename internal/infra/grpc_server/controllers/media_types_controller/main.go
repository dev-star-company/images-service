package media_types_controller

import (
	"context"
	"images-service/generated_protos/media_types_proto"
	"images-service/internal/app/ent"
)

type Controller interface {
	media_types_proto.MediaTypesServiceServer

	Create(ctx context.Context, in *media_types_proto.CreateRequest) (*media_types_proto.CreateResponse, error)
	Get(ctx context.Context, in *media_types_proto.GetRequest) (*media_types_proto.GetResponse, error)
	Update(ctx context.Context, in *media_types_proto.UpdateRequest) (*media_types_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *media_types_proto.DeleteRequest) (*media_types_proto.DeleteResponse, error)
	List(ctx context.Context, in *media_types_proto.ListRequest) (*media_types_proto.ListResponse, error)
}

type controller struct {
	media_types_proto.UnimplementedMediaTypesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
