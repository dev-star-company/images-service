package hosts_controller

import (
	"context"
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos"
)

type Controller interface {
	hosts_proto.HostsServiceServer

	Create(ctx context.Context, in *hosts_proto.CreateRequest) (*hosts_proto.CreateResponse, error)
	Get(ctx context.Context, in *hosts_proto.GetRequest) (*hosts_proto.GetResponse, error)
	Update(ctx context.Context, in *hosts_proto.UpdateRequest) (*hosts_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *hosts_proto.DeleteRequest) (*hosts_proto.DeleteResponse, error)
	List(ctx context.Context, in *hosts_proto.ListRequest) (*hosts_proto.ListResponse, error)
}

type controller struct {
	hosts_proto.UnimplementedHostsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
