package images_controller

import (
	"context"

	"images-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *images_proto.DeleteRequest) (*images_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ImagesNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.Images.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("images", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &images_proto.DeleteResponse{}, nil
}
