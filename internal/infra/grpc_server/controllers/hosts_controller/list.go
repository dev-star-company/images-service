package hosts_controller

import (
	"context"
	"errors"
	"images-service/internal/adapters/grpc_controllers"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/host"
	"images-service/internal/app/ent/schema"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_proto"

	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *host_proto.ListRequest) (*host_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Hosts.Query()

	if in.Name != nil {
		query = query.Where(host.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying host", err)
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
				query = query.Order(ent.Asc(host.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(host.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	host, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying host", err)
	}

	responseHosts := make([]*host_proto.Hosts, len(host))
	for i, acc := range host {
		responseHosts[i] = grpc_controllers.HostsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &host_proto.ListResponse{
		Rows:  responseHosts,
		Count: uint32(count),
	}, nil
}
