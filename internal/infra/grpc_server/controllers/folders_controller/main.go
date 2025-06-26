package folders_controller

import (
	"context"
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"
)

type Controller interface {
	folders_proto.FoldersServiceServer

	Create(ctx context.Context, in *folders_proto.CreateRequest) (*folders_proto.CreateResponse, error)
	Get(ctx context.Context, in *folders_proto.GetRequest) (*folders_proto.GetResponse, error)
	Update(ctx context.Context, in *folders_proto.UpdateRequest) (*folders_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *folders_proto.DeleteRequest) (*folders_proto.DeleteResponse, error)
	List(ctx context.Context, in *folders_proto.ListRequest) (*folders_proto.ListResponse, error)
}

type controller struct {
	folders_proto.UnimplementedFoldersServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
