package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/media_types_proto"
)

func MediaTypesToProto(media_types *ent.MediaTypes) *media_types_proto.MediaTypes {
	if media_types == nil {
		return nil
	}

	cur := &media_types_proto.MediaTypes{
		Id:        uint32(media_types.ID),
		Name:      media_types.Name,
		CreatedAt: media_types.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if media_types.DeletedAt != nil {
		x := media_types.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
