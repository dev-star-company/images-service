package hosts_controller

import (
	"context"
	"errors"
	"images-service/generated_protos/hosts_proto"
	grpc_convertions "images-service/internal/adapters/grpc"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/hosts"
	"images-service/internal/app/ent/schema"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *hosts_proto.ListRequest) (*hosts_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Hosts.Query()

	if in.Name != nil {
		query = query.Where(hosts.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying hosts", err)
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
				query = query.Order(ent.Asc(hosts.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(hosts.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	hosts, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying hosts", err)
	}

	responseHosts := make([]*hosts_proto.Hosts, len(hosts))
	for i, acc := range hosts {
		responseHosts[i] = grpc_convertions.HostsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &hosts_proto.ListResponse{
		Rows:  responseHosts,
		Count: uint32(count),
	}, nil
}
