package hosts_controller

import (
	"context"
	"images-service/generated_protos/hosts_proto"
	"images-service/internal/app/ent"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *hosts_proto.UpdateRequest) (*hosts_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var hosts *ent.Hosts

	hostsQ := tx.Hosts.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		hostsQ.SetName(string(*in.Name))
	}

	hostsQ.SetUpdatedBy(int(in.RequesterId))

	hosts, err = hostsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("hosts", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &hosts_proto.UpdateResponse{
		RequesterId: uint32(hosts.CreatedBy),
		Name:        string(*hosts.Name),
	}, nil
}
