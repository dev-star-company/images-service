package host_urls_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/host_urls"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *host_urls_proto.GetRequest) (*host_urls_proto.GetResponse, error) {
	host_urls, err := c.Db.HostURLS.
		Query().
		Where(host_urls.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.HostURLSNotFound(int(in.Id))
	}

	return &host_urls_proto.GetResponse{
		RequesterId: uint32(host_urls.CreatedBy),
		Default:     bool(host_urls.Default),
		Url:         string(host_urls.Url),
		Name:        string(*host_urls.Name),
	}, nil
}
