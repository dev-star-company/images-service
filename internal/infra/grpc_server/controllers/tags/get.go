package tags_controller

import (
	"context"
	"github.com/dev-star-company/protos-go/images_service/generated_protos/tags_proto"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/tags"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *tags_proto.GetRequest) (*tags_proto.GetResponse, error) {
	tags, err := c.Db.Tags.
		Query().
		Where(tags.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.TagsNotFound(int(in.Id))
	}

	return &tags_proto.GetResponse{
		Name:        tags.Name,
	}, nil
}
