package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/database"
)

func UnfollowUser(follower string, followee string) error {
	db := database.DB

	followerUUID := pgtype.UUID{}
	followerUUID.Scan(follower)
	followeeUUID := pgtype.UUID{}
	followeeUUID.Scan(followee)

	params := database.DeleteFollowRelParams{
		Follower: followerUUID,
		Followee: followeeUUID,
	}

	err := db.DeleteFollowRel(context.Background(), params)
	if err != nil {
		return err
	}

	return nil
}
