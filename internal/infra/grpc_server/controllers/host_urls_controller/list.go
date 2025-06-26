package host_urls_controller

import (
	"context"
	"errors"
	grpc_controllers "images-service/internal/adapters/grpc"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/host_urls"
	"images-service/internal/app/ent/schema"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *host_urls_proto.ListRequest) (*host_urls_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.HostURLS.Query()

	if in.Name != nil {
		query = query.Where(host_urls.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying host_urls", err)
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
				query = query.Order(ent.Asc(host_urls.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(host_urls.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	host_urls, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying host_urls", err)
	}

	responseHostURLS := make([]*host_urls_proto.HostURLS, len(host_urls))
	for i, acc := range host_urls {
		responseHostURLS[i] = grpc_controllers.HostURLSToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_urls_proto.ListResponse{
		Rows:  responseHostURLS,
		Count: uint32(count),
	}, nil
}
