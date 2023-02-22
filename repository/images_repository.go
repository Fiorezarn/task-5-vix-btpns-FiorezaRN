package repository

import (
	"BTPN-API-go/model/entity"
	"context"
	"database/sql"
)

type ImageRepository interface {
	Create(ctx context.Context, tx *sql.Tx, image entity.Images) entity.Images
	Delete(ctx context.Context, tx *sql.Tx, image entity.Images)
	FindById(ctx context.Context, tx *sql.Tx, imageId string) (entity.Images, error)
}
