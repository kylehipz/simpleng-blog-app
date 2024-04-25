package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/database"
)

func FollowUser(follower string, followee string) (*database.Follow, error) {
	db := database.DB

	followerUUID := pgtype.UUID{}
	followerUUID.Scan(follower)
	followeeUUID := pgtype.UUID{}
	followeeUUID.Scan(followee)

	params := database.CreateFollowRelParams{
		Follower: followerUUID,
		Followee: followeeUUID,
	}

	result, err := db.CreateFollowRel(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
