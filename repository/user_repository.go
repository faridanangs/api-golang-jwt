package repository

import (
	"context"
	"database/sql"
	"user_service/models/entity"
)

type UserRepositry interface {
	Save(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Delete(ctx context.Context, tx *sql.Tx, user entity.Users)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Users
	FindById(ctx context.Context, tx *sql.Tx, userId string) (entity.Users, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error)
}
