package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/kylehipz/simpleng-blog-app/blog-post-service/types"
)

type PostgresStore struct {
	DB *sql.DB
}

func (ps *PostgresStore) FindUsers(id string) (*types.User, error) {
	var blogPost *types.User

	rows, err := ps.DB.Query("SELECT * FROM post WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rows.Scan(&blogPost)

	return blogPost, nil
}
