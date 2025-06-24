package tags_controller

import (
	"context"
	"images-service/generated_protos/tags_proto"
	"images-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *tags_proto.DeleteRequest) (*tags_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.TagsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.Tags.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("tags", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &tags_proto.DeleteResponse{}, nil
}
