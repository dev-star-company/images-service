package hosts_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/host"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *host_proto.GetRequest) (*host_proto.GetResponse, error) {
	host, err := c.Db.Hosts.
		Query().
		Where(host.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.HostsNotFound(int(in.Id))
	}

	return &host_proto.GetResponse{
		RequesterId: uint32(host.CreatedBy),
		Name:        *host.Name,
	}, nil
}
