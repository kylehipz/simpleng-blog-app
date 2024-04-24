package database

import (
	"context"

	"github.com/jackc/pgx/v5"

	db "simpleng-blog-app/common/internal/database"
)

var DB *db.Queries

func Connect() error {
	dsn := "host=localhost port=5432 user=postgres password=postgres sslmode=disable database=simpleng-blog-app"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	DB = db.New(conn)

	return nil
}
