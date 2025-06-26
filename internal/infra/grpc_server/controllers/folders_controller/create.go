package folders_controller

import (
	"context"
	"images-service/internal/app/ent/folders"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *folders_proto.CreateRequest) (*folders_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	create, err := c.Db.Folders.Create().
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

	return &folders_proto.CreateResponse{
		RequesterId: uint32(create.CreatedBy),
		FolderId:    uint32(folders.FolderId),
		Name:        string(*folders.Name),
	}, nil
}
