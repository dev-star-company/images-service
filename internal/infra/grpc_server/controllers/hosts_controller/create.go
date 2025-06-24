package hosts_controller

import (
	"context"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/protos"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *hosts_proto.CreateRequest) (*hosts_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	create, err := c.Db.Hosts.Create().
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

	return &hosts_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		Name:        string(*create.Name),
	}, nil
}
