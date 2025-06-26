package host_urls_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *host_urls_proto.UpdateRequest) (*host_urls_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var host_urls *ent.HostURLS

	host_urlsQ := tx.HostURLS.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		host_urlsQ.SetName(string(*in.Name))
	}

	host_urlsQ.SetUpdatedBy(int(in.RequesterId))

	host_urls, err = host_urlsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
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
		RequesterId: uint32(host_urls.CreatedBy),
		Default:     bool(host_urls.Default),
		Url:         string(host_urls.Url),
		Name:        string(*host_urls.Name),
	}, nil
}
