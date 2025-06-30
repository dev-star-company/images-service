package host_urls_controller

import (
	"context"
	"images-service/internal/infra/grpc_server/controllers"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *host_urls_proto.CreateRequest) (*host_urls_proto.CreateResponse, error) {

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

	_, err = c.Db.HostURLS.Create().
		SetCreatedBy(requester.ID).
		SetUpdatedBy(requester.ID).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("host_urls", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_urls_proto.CreateResponse{
		RequesterUuid: in.RequesterUuid,
		Default:       bool(in.Default),
		Url:           string(in.Url),
		Name:          string(in.Name),
	}, nil
}
