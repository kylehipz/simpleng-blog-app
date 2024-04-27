package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/database"
)

func UpdateBlog(id string, title string, body string) (*database.Blog, error) {
	db := database.DB
	uuid := pgtype.UUID{}
	uuid.Scan(id)

	params := database.UpdateBlogParams{
		ID:    uuid,
		Title: title,
		Body:  body,
	}

	updatedBlog, err := db.UpdateBlog(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &updatedBlog, nil
}
