package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/database"
)

func GetHomeFeed(follower string, limit int32, page int32) (int64, []database.Blog, error) {
	db := database.DB

	followerUUID := pgtype.UUID{}
	followerUUID.Scan(follower)

	offset := int32(0)

	if page > 0 {
		offset = (page - 1) * limit
	}

	params := database.GetBlogsFromFolloweesParams{
		Follower: followerUUID,
		Limit:    limit,
		Offset:   offset,
	}

	count, err := db.GetBlogsFromFolloweesCount(context.Background(), followerUUID)
	if err != nil {
		return 0, nil, err
	}

	blogs, err := db.GetBlogsFromFollowees(context.Background(), params)
	if err != nil {
		return 0, nil, err
	}

	return count, blogs, nil
}
