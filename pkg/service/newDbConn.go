package service

import (
	"database/sql"
	"log"
	"time"
)

func GetDbClient(l *log.Logger) *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/banking")
	if err != nil {
		l.Println("Error opening database conn pool : ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
