package folders_controller

import (
	"context"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *folders_proto.CreateRequest) (*folders_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	defer tx.Rollback()

	_, err = c.Db.Folders.Create().
		SetName(in.Name).
		SetHostUrlsID(uint32(in.HostUrlsId)).
		SetFolderID(uint32(in.FolderId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("folders", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &folders_proto.CreateResponse{
		HostUrlsId: in.HostUrlsId,
		FolderId:   in.FolderId,
	}, nil
}
