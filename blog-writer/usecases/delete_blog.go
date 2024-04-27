package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/database"
)

func DeleteBlog(id string) error {
	db := database.DB
	uuid := pgtype.UUID{}
	uuid.Scan(id)

	err := db.DeleteBlog(context.Background(), uuid)
	if err != nil {
		return err
	}

	return nil
}
