package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/folders_proto"
)

func FoldersToProto(folders *ent.Folders) *folders_proto.Folders {
	if folders == nil {
		return nil
	}

	cur := &folders_proto.Folders{
		Id:         uint32(folders.ID),
		Name:       folders.Name,
		FolderId:   uint32(*folders.FolderID),
		HostUrlsId: uint32(*folders.HostUrlsID),
		CreatedAt:  folders.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if folders.DeletedAt != nil {
		x := folders.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
