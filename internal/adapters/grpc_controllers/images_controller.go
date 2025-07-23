package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
)

func ImagesToProto(hosts *ent.Images) *images_proto.Images {
	if hosts == nil {
		return nil
	}

	cur := &images_proto.Images{
		Id:          uint32(hosts.ID),
		Name:        hosts.Name,
		Uuid:        hosts.UUID,
		FolderId:    uint32(*hosts.FolderID),
		MediaTypeId: uint32(*hosts.MediaTypeID),
		CreatedAt:   hosts.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if hosts.DeletedAt != nil {
		x := hosts.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
