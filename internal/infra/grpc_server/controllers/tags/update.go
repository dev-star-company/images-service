package tags_controller

import (
	"context"
	"images-service/generated_protos/tags_proto"
	"images-service/internal/app/ent"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *tags_proto.UpdateRequest) (*tags_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var tags *ent.Tags

	tagsQ := tx.Tags.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		tagsQ.SetName(string(*in.Name))
	}

	tagsQ.SetUpdatedBy(int(in.RequesterId))

	tags, err = tagsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("tags", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &tags_proto.UpdateResponse{
		RequesterId: uint32(tags.CreatedBy),
		Name:        string(*tags.Name),
	}, nil
}
