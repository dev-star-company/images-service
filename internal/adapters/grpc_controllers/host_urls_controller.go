package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_urls_proto"
)

func HostURLSToProto(host_urls *ent.HostURLS) *host_urls_proto.HostURLs {
	if host_urls == nil {
		return nil
	}

	cur := &host_urls_proto.HostURLs{
		Id:        uint32(host_urls.ID),
		Name:      host_urls.Name,
		Url:       host_urls.URL,
		Default:   bool(host_urls.Default),
		CreatedAt: host_urls.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if host_urls.DeletedAt != nil {
		x := host_urls.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
