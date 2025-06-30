package controllers

import (
	"context"
	"fmt"
	"images-service/internal/app/ent"
	"images-service/internal/app/ent/images"
	"images-service/internal/pkg/utils/parser"

	"github.com/dev-star-company/service-errors/errs"
)

func GetUserFromUuid(tx *ent.Tx, ctx context.Context, requesterUuid string) (*ent.Images, error) {
	if requesterUuid == "" {
		return nil, errs.RequesterIDRequired()
	}

	uuidRequester, err := parser.Uuid(requesterUuid)
	if err != nil {
		return nil, fmt.Errorf("invalid requester UUID: %w", err)
	}

	requesterUser, err := tx.Images.Query().
		Where(images.UUIDEQ(uuidRequester)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching requester user: %w", err)
	}
	if requesterUser == nil {
		return nil, errs.UserNotFound(0)
	}
	return requesterUser, nil
}
