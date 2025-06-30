package folders_controller

import (
	"context"
	"errors"
	"images-service/internal/adapters/grpc_controllers"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/folders"
	"images-service/internal/app/ent/schema"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) List(ctx context.Context, in *folders_proto.ListRequest) (*folders_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Folders.Query()

	if in.Name != "" {
		query = query.Where(folders.Name(string(in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying folders", err)
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
				query = query.Order(ent.Asc(folders.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(folders.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	folders, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying folders", err)
	}

	responseFolders := make([]*folders_proto.Folders, len(folders))
	for i, acc := range folders {
		responseFolders[i] = grpc_controllers.FoldersToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &folders_proto.ListResponse{
		Rows:  responseFolders,
		Count: uint32(count),
	}, nil
}
