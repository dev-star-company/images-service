package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
)

func ImagesToProto(image *ent.Images) *images_proto.Images {
	if image == nil {
		return nil
	}

	cur := &images_proto.Images{
		Id:           uint32(image.ID),
		Name:         image.Name,
		CloudflareId: image.CloudflareID,
		Url:          image.URL,
		Size:         image.Size,
		ContentType:  image.ContentType,
		FolderId:     uint32(*image.FolderID),
		MediaTypeId:  uint32(*image.MediaTypeID),
		CreatedAt:    image.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if image.DeletedAt != nil {
		x := image.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
