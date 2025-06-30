package folders_controller

import (
	"context"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/folders"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Get(ctx context.Context, in *folders_proto.GetRequest) (*folders_proto.GetResponse, error) {
	folders, err := c.Db.Folders.
		Query().
		Where(folders.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.FoldersNotFound(int(in.Id))
	}

	return &folders_proto.GetResponse{
		FolderId: uint32(*folders.FolderID),
		Name:     string(folders.Name),
	}, nil
}
