package service

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/url"
	"time"
)

const (
	scheme   = "postgres"
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "1234"
	dbname   = "banking"
)

func GetDbClient() *sql.DB {
	dsn := url.URL{
		Scheme: scheme,
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   dbname,
	}

	q := dsn.Query()
	q.Set("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		panic(err)
	}

	log.Printf("successfully connected to database %s", dsn.String())

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
