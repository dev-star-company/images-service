package images_controller

import (
	"context"
	"images-service/generated_protos/images_proto"
	"images-service/internal/app/ent"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Update(ctx context.Context, in *images_proto.UpdateRequest) (*images_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	var images *ent.Images

	imagesQ := tx.Images.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		imagesQ.SetName(string(*in.Name))
	}

	imagesQ.SetUpdatedBy(int(in.RequesterId))

	images, err = imagesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("images", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &images_proto.UpdateResponse{
		RequesterId: uint32(images.CreatedBy),
		Name:        string(*images.Name),
		Uuid:        string(*images.Uuid),
		FolderId:    uint32(images.FolderId),
		MediaTypeId: uint32(images.MediaTypeId),
	}, nil
}
