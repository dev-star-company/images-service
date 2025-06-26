package images_controller

import (
	"context"
	"images-service/internal/app/ent/images"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *images_proto.CreateRequest) (*images_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	create, err := c.Db.Images.Create().
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

	return &images_proto.CreateResponse{
		RequesterId: uint32(images.CreatedBy),
		Name:        string(*images.Name),
		Uuid:        string(*images.Uuid),
		FolderId:    uint32(images.FolderId),
		MediaTypeId: uint32(images.MediaTypeId),
	}, nil
}
