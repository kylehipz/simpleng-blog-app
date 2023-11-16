package store

import (
	"database/sql"
	"fmt"

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
	existingFollow := &types.Follow{}

	findRow := ps.DB.QueryRow(
		"SELECT * FROM follow WHERE follower = $1 AND followee = $2",
		follow.Follower,
		follow.Followee,
	)

	findRow.Scan(&existingFollow.Follower, &existingFollow.Followee)

	if existingFollow.Follower == follow.Follower && existingFollow.Followee == follow.Followee {
		return nil, fmt.Errorf("User already followed")
	}

	insertResult := &types.Follow{}

	row := ps.DB.QueryRow(
		"INSERT INTO follow VALUES($1, $2) RETURNING *",
		follow.Follower,
		follow.Followee,
	)

	row.Scan(&insertResult.Follower, &insertResult.Followee)

	return insertResult, nil
}
