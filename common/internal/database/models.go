// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Follow struct {
	ID        pgtype.UUID
	Follower  pgtype.UUID
	Followee  pgtype.UUID
	CreatedAt pgtype.Timestamptz
}

type User struct {
	ID        pgtype.UUID
	UserName  string
	CreatedAt pgtype.Timestamptz
}
