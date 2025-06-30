package host_urls_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/infra/grpc_server/controllers"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *host_urls_proto.UpdateRequest) (*host_urls_proto.UpdateResponse, error) {
	if in.RequesterUuid == "" {
		return nil, errs.HostURLSNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}
	requester, err := controllers.GetUserFromUuid(tx, ctx, in.RequesterUuid)
	if err != nil {
		return nil, err
	}

	host_urlsQ := tx.HostURLS.UpdateOneID(int(in.Id))

	host_urlsQ.SetUpdatedBy(requester.ID)

	_, err = host_urlsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.HostURLSNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("host_urls", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_urls_proto.UpdateResponse{
		RequesterUuid: in.RequesterUuid,
	}, nil
}
