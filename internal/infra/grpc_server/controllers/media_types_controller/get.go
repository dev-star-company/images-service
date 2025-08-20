package media_types_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/mediatypes"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/media_types_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *media_types_proto.GetRequest) (*media_types_proto.GetResponse, error) {
	media_types, err := c.Db.MediaTypes.
		Query().
		Where(mediatypes.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.MediaTypesNotFound(int(in.Id))
	}

	return &media_types_proto.GetResponse{
		Name: media_types.Name,
	}, nil
}
