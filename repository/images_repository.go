package repository

import (
	"context"
	"database/sql"
	"user_service/models/entity"
)

type ImageRepository interface {
	Save(ctx context.Context, tx *sql.Tx, request entity.Images) entity.Images
	Delete(ctx context.Context, tx *sql.Tx, request entity.Images)
	FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Images, error)
}
