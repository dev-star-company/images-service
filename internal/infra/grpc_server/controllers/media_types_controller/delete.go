package media_types_controller

import (
	"context"
	"images-service/generated_protos/media_types_proto"
	"images-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *media_types_proto.DeleteRequest) (*media_types_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.MediaTypesNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.MediaTypes.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("media_types", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &media_types_proto.DeleteResponse{}, nil
}
