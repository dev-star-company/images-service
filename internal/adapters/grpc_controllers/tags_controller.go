package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/tags_proto"
)

func TagsToProto(tags *ent.Tags) *tags_proto.Tags {
	if tags == nil {
		return nil
	}

	cur := &tags_proto.Tags{
		Id:        uint32(tags.ID),
		Name:      *tags.Name,
		CreatedBy: uint32(tags.CreatedBy),
		UpdatedBy: uint32(tags.UpdatedBy),
		CreatedAt: tags.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: tags.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if tags.DeletedAt != nil {
		x := tags.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if tags.DeletedBy != nil {
		x := uint32(*tags.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
