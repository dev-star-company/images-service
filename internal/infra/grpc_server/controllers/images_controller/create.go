package images_controller

import (
	"context"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *images_proto.CreateRequest) (*images_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}
	defer tx.Rollback()

	_, err = c.Db.Images.Create().
		SetName(in.Name).
		SetMediaTypeID(int(in.MediaTypeId)).
		SetFolderID(int(in.FolderId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &images_proto.CreateResponse{}, nil
}
