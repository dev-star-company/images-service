package host_urls_controller

import (
	"context"
	"images-service/internal/app/ent/host_urls"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *host_urls_proto.CreateRequest) (*host_urls_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	create, err := c.Db.HostURLS.Create().
		SetName(in.Name).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_urls_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		Default:     bool(host_urls.Default),
		Url:         string(host_urls.Url),
		Name:        string(*host_urls.Name),
	}, nil
}
