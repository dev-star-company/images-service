package tags_controller

import (
	"context"
	"errors"
	"images-service/generated_protos/tags_proto"
	grpc_controllers "images-service/internal/adapters/grpc"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/schema"
	"images-service/internal/app/ent/tags"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *tags_proto.ListRequest) (*tags_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Tags.Query()

	if in.Name != nil {
		query = query.Where(tags.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying tags", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	if in.Orderby != nil {
		if in.Orderby.Id != nil {
			switch *in.Orderby.Id {
			case "ASC":
				query = query.Order(ent.Asc(tags.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(tags.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	tags, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying tags", err)
	}

	responseTags := make([]*tags_proto.Tags, len(tags))
	for i, acc := range tags {
		responseTags[i] = grpc_controllers.TagsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &tags_proto.ListResponse{
		Rows:  responseTags,
		Count: uint32(count),
	}, nil
}
