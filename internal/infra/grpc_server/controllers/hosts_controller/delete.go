package hosts_controller

import (
	"context"

	"github.com/dev-star-company/protos-go/images_service/protos/hosts"

	"images-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *hosts_proto.DeleteRequest) (*hosts_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.HostsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	err = tx.Hosts.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("hosts", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &hosts_proto.DeleteResponse{}, nil
}
