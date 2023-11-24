package app

import (
	"database/sql"
	"time"
	"user_service/helper"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/yt_users_service?parseTime=true")
	helper.PanicError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)

	return db
}
