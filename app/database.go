package app

import (
	"akmmp241/belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "joko:akuanakhebat@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
