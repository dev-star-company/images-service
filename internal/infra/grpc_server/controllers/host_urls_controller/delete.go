package host_urls_controller

import (
	"context"

	"images-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *host_urls_proto.DeleteRequest) (*host_urls_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.HostURLSNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.HostURLS.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("host_urls", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_urls_proto.DeleteResponse{}, nil
}
