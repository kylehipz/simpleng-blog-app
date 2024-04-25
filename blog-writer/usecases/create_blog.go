package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/database"
)

func CreateBlog(author string, title string, body string) (*database.Blog, error) {
	authorUUID := pgtype.UUID{}
	authorUUID.Scan(author)

	params := database.CreateBlogParams{
		Author: authorUUID,
		Title:  title,
		Body:   body,
	}

	newBlog, err := database.DB.CreateBlog(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &newBlog, nil
}
