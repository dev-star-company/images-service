package tags_controller

import (
	"context"
	"images-service/generated_protos/tags_proto"
	"images-service/internal/app/ent"
)

type Controller interface {
	tags_proto.TagsServiceServer

	Create(ctx context.Context, in *tags_proto.CreateRequest) (*tags_proto.CreateResponse, error)
	Get(ctx context.Context, in *tags_proto.GetRequest) (*tags_proto.GetResponse, error)
	Update(ctx context.Context, in *tags_proto.UpdateRequest) (*tags_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *tags_proto.DeleteRequest) (*tags_proto.DeleteResponse, error)
	List(ctx context.Context, in *tags_proto.ListRequest) (*tags_proto.ListResponse, error)
}

type controller struct {
	tags_proto.UnimplementedTagsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
