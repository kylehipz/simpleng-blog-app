package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	db "simpleng-blog-app/common/internal/database"
)

var DB *db.Queries

func Connect() error {
	dsnFormat := "host=%s port=%s user=%s password=%s database=%s"

	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(dsnFormat, host, dbPort, user, password, dbName)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}

	DB = db.New(conn)

	return nil
}
