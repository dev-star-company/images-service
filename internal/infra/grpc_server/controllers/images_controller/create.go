package images_controller

import (
	"context"
	"images-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/images_service/generated_protos/images_proto"
	"github.com/dev-star-company/service-errors/errs"
)

func (c *controller) Create(ctx context.Context, in *images_proto.CreateRequest) (*images_proto.CreateResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartTransactionError(err)
	}

	// Upload para Cloudflare Images
	metadata := map[string]string{
		"name":      in.Name,
		"folder_id": string(rune(in.FolderId)),
	}

	uploadResp, err := c.CloudflareClient.UploadImage(ctx, in.ImageData, in.Name, metadata)
	if err != nil {
		return nil, utils.Rollback(tx, errs.CreateError("cloudflare_upload", err))
	}

	// Gerar URL pública
	publicURL := c.CloudflareClient.GetImageURL(uploadResp.Result.ID, "public")

	// Salvar no banco
	image, err := tx.Images.Create().
		SetName(in.Name).
		SetCloudflareID(uploadResp.Result.ID).
		SetURL(publicURL).
		SetSize(int64(len(in.ImageData))).
		SetContentType(in.ContentType).
		SetFolderID(int(in.FolderId)).
		Save(ctx)

	if err != nil {
		// Se falhar ao salvar no banco, deletar da Cloudflare
		c.CloudflareClient.DeleteImage(ctx, uploadResp.Result.ID)
		return nil, utils.Rollback(tx, errs.CreateError("images", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitTransactionError(err))
	}

	return &images_proto.CreateResponse{
		Id:           uint32(image.ID),
		CloudflareId: uploadResp.Result.ID,
		Url:          publicURL,
	}, nil
}
