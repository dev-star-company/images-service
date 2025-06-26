package host_urls_controller

import (
	"context"
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
)

type Controller interface {
	host_urls_proto.HostURLsServiceServer

	Create(ctx context.Context, in *host_urls_proto.CreateRequest) (*host_urls_proto.CreateResponse, error)
	Get(ctx context.Context, in *host_urls_proto.GetRequest) (*host_urls_proto.GetResponse, error)
	Update(ctx context.Context, in *host_urls_proto.UpdateRequest) (*host_urls_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *host_urls_proto.DeleteRequest) (*host_urls_proto.DeleteResponse, error)
	List(ctx context.Context, in *host_urls_proto.ListRequest) (*host_urls_proto.ListResponse, error)
}

type controller struct {
	host_urls_proto.UnimplementedHostURLsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
