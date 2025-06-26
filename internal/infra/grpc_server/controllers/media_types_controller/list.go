package media_types_controller

import (
	"context"
	"errors"
	"images-service/generated_protos/media_types_proto"
	grpc_controllers "images-service/internal/adapters/grpc"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/media_types"
	"images-service/internal/app/ent/schema"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *media_types_proto.ListRequest) (*media_types_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.MediaTypes.Query()

	if in.Name != nil {
		query = query.Where(media_types.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying media_types", err)
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
				query = query.Order(ent.Asc(media_types.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(media_types.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	media_types, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying media_types", err)
	}

	responseMediaTypes := make([]*media_types_proto.MediaTypes, len(media_types))
	for i, acc := range media_types {
		responseMediaTypes[i] = grpc_controllers.MediaTypesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &media_types_proto.ListResponse{
		Rows:  responseMediaTypes,
		Count: uint32(count),
	}, nil
}
