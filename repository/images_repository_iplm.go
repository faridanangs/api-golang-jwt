package repository

import (
	"context"
	"database/sql"
	"errors"
	"user_service/helper"
	"user_service/models/entity"
)

type ImagesRepositoryIplm struct{}

func NewImagesRepositoryIplm() ImageRepository {
	return &ImagesRepositoryIplm{}
}

func (repository *ImagesRepositoryIplm) Save(ctx context.Context, tx *sql.Tx, request entity.Images) entity.Images {
	SQL := "insert into images(id, path, created_at, updated_at) values(?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, request.Id, request.Path, request.CreatedAt, request.UpdatedAt)
	helper.PanicError(err)

	return request

}

func (repository *ImagesRepositoryIplm) Delete(ctx context.Context, tx *sql.Tx, request entity.Images) {
	SQL := "DELETE FROM images where id = ?"
	_, err := tx.ExecContext(ctx, SQL, request.Id)
	helper.PanicError(err)
}
func (repository *ImagesRepositoryIplm) FindById(ctx context.Context, tx *sql.Tx, requestId string) (entity.Images, error) {
	SQL := "select * from images where id = ?"
	row, err := tx.QueryContext(ctx, SQL, requestId)
	helper.PanicError(err)

	var image entity.Images
	if row.Next() {
		err := row.Scan(&image.Id, &image.Path, &image.CreatedAt, &image.UpdatedAt)
		helper.PanicError(err)
		return image, nil
	} else {
		return image, errors.New("image not found")
	}
}
