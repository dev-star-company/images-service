package grpc_controllers

import (
	"images-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/host_proto"
)

func HostsToProto(hosts *ent.Hosts) *host_proto.Host {
	if hosts == nil {
		return nil
	}

	cur := &host_proto.Host{
		Id:        uint32(hosts.ID),
		Name:      hosts.Name,
		CreatedBy: uint32(hosts.CreatedBy),
		UpdatedBy: uint32(hosts.UpdatedBy),
		CreatedAt: hosts.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: hosts.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if hosts.DeletedAt != nil {
		x := hosts.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if hosts.DeletedBy != nil {
		x := uint32(*hosts.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
