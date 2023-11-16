package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/kylehipz/simpleng-blog-app/users-service/types"
)

type PostgresStore struct {
	DB *sql.DB
}

func (ps *PostgresStore) FindUsers() ([]*types.User, error) {
	var users []*types.User

	rows, err := ps.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &types.User{}

		rows.Scan(&user.ID, &user.Email, &user.UserName)

		users = append(users, user)
	}

	return users, nil
}

func (ps *PostgresStore) InsertFollow(follow *types.Follow) (*types.Follow, error) {
	return nil, nil
}
