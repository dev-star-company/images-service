package folders_controller

import (
	"context"
	"images-service/internal/infra/grpc_server/controllers"
	"images-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Delete(ctx context.Context, in *folders_proto.DeleteRequest) (*folders_proto.DeleteResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.FoldersNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	err = tx.Folders.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(requester.ID).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("folders", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &folders_proto.DeleteResponse{}, nil
}
