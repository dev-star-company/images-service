package hosts_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/hosts"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *hosts_proto.GetRequest) (*hosts_proto.GetResponse, error) {
	hosts, err := c.Db.Hosts.
		Query().
		Where(hosts.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.HostsNotFound(int(in.Id))
	}

	return &hosts_proto.GetResponse{
		RequesterId: uint32(hosts.CreatedBy),
		Name:        *hosts.Name,
	}, nil
}
