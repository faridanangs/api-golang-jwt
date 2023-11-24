package repository

import (
	"context"
	"database/sql"
	"errors"
	"user_service/helper"
	"user_service/models/entity"
)

type UserRepositoryIplm struct{}

func NewRepositoryIplm() UserRepositry {
	return &UserRepositoryIplm{}
}

func (repository *UserRepositoryIplm) Save(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "insert into users(id, first_name, last_name, email, password, created_at, updated_at) values (?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	helper.PanicError(err)
	return user
}
func (repository *UserRepositoryIplm) Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users {
	SQL := "update users set first_name=?, last_name=?, updated_at=? where id=?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
		user.Id,
	)
	helper.PanicError(err)

	return user
}
func (repository *UserRepositoryIplm) Delete(ctx context.Context, tx *sql.Tx, user entity.Users) {
	SQL := "delete from users where id=?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
	)
	helper.PanicError(err)
}
func (repository *UserRepositoryIplm) FindById(ctx context.Context, tx *sql.Tx, userId string) (entity.Users, error) {
	SQL := "select id, first_name, last_name, email, password, created_at, updated_at from users where id=?"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		userId,
	)
	helper.PanicError(err)
	defer rows.Close()

	var user entity.Users
	if rows.Next() {
		err := rows.Scan(
			&user.Id, &user.FirstName, &user.LastName,
			&user.Email, &user.Password, &user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicError(err)
		return user, nil
	} else {
		return user, errors.New("user not found in db")
	}

}
func (repository *UserRepositoryIplm) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error) {
	SQL := "select id, first_name, last_name, email, password, created_at, updated_at from users where email=?"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		email,
	)
	helper.PanicError(err)
	defer rows.Close()

	var user entity.Users
	if rows.Next() {
		err := rows.Scan(
			&user.Id, &user.FirstName, &user.LastName,
			&user.Email, &user.Password, &user.CreatedAt,
			&user.UpdatedAt,
		)
		helper.PanicError(err)
		return user, nil
	} else {
		return user, errors.New("email not found")
	}

}
func (repository *UserRepositoryIplm) FindAll(ctx context.Context, tx *sql.Tx) []entity.Users {
	SQL := "select id, first_name, last_name, email, password, created_at, updated_at from users"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)
	helper.PanicError(err)
	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		var user entity.Users
		err := rows.Scan(
			&user.Id, &user.FirstName, &user.LastName,
			&user.Email, &user.Password, &user.CreatedAt,
			&user.UpdatedAt,
		)
		users = append(users, user)
		helper.PanicError(err)
	}

	return users
}
