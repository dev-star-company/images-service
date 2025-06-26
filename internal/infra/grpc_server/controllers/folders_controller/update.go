package folders_controller

import (
	"context"
	"images-service/generated_protos/folders_proto"
	"images-service/internal/app/ent"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *folders_proto.UpdateRequest) (*folders_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var folders *ent.Folders

	foldersQ := tx.Folders.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		foldersQ.SetName(string(*in.Name))
	}

	foldersQ.SetUpdatedBy(int(in.RequesterId))

	folders, err = foldersQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
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
		RequesterId: uint32(folders.CreatedBy),
		FolderId:    uint32(folders.FolderId),
		Name:        string(*folders.Name),
	}, nil
}
