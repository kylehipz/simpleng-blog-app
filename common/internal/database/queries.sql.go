// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package database

import "context"

const createUser = `-- name: CreateUser :one
INSERT INTO users (user_name) VALUES ($1) RETURNING id, user_name, created_at
`

func (q *Queries) CreateUser(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRow(ctx, createUser, userName)
	var i User
	err := row.Scan(&i.ID, &i.UserName, &i.CreatedAt)
	return i, err
}
