package folders_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/infra/grpc_server/controllers"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *folders_proto.UpdateRequest) (*folders_proto.UpdateResponse, error) {
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

	foldersQ := tx.Folders.UpdateOneID(int(in.Id))

	foldersQ.SetUpdatedBy(requester.ID)

	_, err = foldersQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.FoldersNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("folders", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &folders_proto.UpdateResponse{
		RequesterUuid: in.RequesterUuid,
	}, nil
}
