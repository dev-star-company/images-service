package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
)

func HostURLSToProto(host_urls *ent.HostURLS) *host_urls_proto.HostURLS {
	if host_urls == nil {
		return nil
	}

	cur := &host_urls_proto.HostURLS{
		Id:        uint32(host_urls.ID),
		Name:      *host_urls.Name,
		Url:       *host_urls.Url,
		Default:   bool(*host_urls.Default),
		CreatedBy: uint32(host_urls.CreatedBy),
		UpdatedBy: uint32(host_urls.UpdatedBy),
		CreatedAt: host_urls.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: host_urls.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if host_urls.DeletedAt != nil {
		x := host_urls.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if host_urls.DeletedBy != nil {
		x := uint32(*host_urls.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
