package folders_controller

import (
	"context"
	"images-service/internal/infra/grpc_server/controllers"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *folders_proto.CreateRequest) (*folders_proto.CreateResponse, error) {

	if in.RequesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	defer tx.Rollback()

	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	_, err = c.Db.Folders.Create().
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("folders", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &folders_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
		FolderId:      uint32(*&in.FolderId),
	}, nil
}
