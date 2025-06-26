package images_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/images"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *images_proto.GetRequest) (*images_proto.GetResponse, error) {
	images, err := c.Db.Images.
		Query().
		Where(images.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ImagesNotFound(int(in.Id))
	}

	return &images_proto.GetResponse{
		RequesterId: uint32(images.CreatedBy),
		Name:        string(*images.Name),
		Uuid:        string(*images.Uuid),
		FolderId:    uint32(images.FolderId),
		MediaTypeId: uint32(images.MediaTypeId),
	}, nil
}
