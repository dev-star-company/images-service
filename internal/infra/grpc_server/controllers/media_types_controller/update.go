package media_types_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/media_types_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *media_types_proto.UpdateRequest) (*media_types_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var media_types *ent.MediaTypes

	media_typesQ := tx.MediaTypes.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		media_typesQ.SetName(string(*in.Name))
	}

	media_types, err = media_typesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("media_types", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &media_types_proto.UpdateResponse{
		Name: string(media_types.Name),
	}, nil
}
