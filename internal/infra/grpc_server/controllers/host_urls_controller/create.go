package host_urls_controller

import (
	"context"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *host_urls_proto.CreateRequest) (*host_urls_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	defer tx.Rollback()

	_, err = c.Db.HostURLS.Create().
		SetDefault(in.Default).
		SetURL(in.Url).
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("host_urls", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_urls_proto.CreateResponse{
		Default: bool(in.Default),
		Url:     string(in.Url),
		Name:    string(in.Name),
	}, nil
}
