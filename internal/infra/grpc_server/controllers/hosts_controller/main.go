package hosts_controller

import (
	"context"
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_proto"
)

type Controller interface {
	host_proto.HostServiceServer

	Create(ctx context.Context, in *host_proto.CreateRequest) (*host_proto.CreateResponse, error)
	Get(ctx context.Context, in *host_proto.GetRequest) (*host_proto.GetResponse, error)
	Update(ctx context.Context, in *host_proto.UpdateRequest) (*host_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *host_proto.DeleteRequest) (*host_proto.DeleteResponse, error)
	List(ctx context.Context, in *host_proto.ListRequest) (*host_proto.ListResponse, error)
}

type controller struct {
	host_proto.UnimplementedHostServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
