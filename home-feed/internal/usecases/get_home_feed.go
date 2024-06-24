package usecases

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"simpleng-blog-app/common/pkg/cache"
	"simpleng-blog-app/common/pkg/database"
)

func GetHomeFeed(follower string, limit int32, page int32) (int64, []database.Blog, error) {
	ctx := context.Background()
	db := database.DB
	count := int64(0)

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

	countStr, err := cache.Cache.HGet(ctx, follower, "count").Result()
	if err != nil {
		fmt.Println("err", err)
		count, err = db.GetBlogsFromFolloweesCount(ctx, followerUUID)
		// Log error
		if err != nil {
			return 0, nil, err
		}

		cache.Cache.HSet(ctx, follower, "count", count)
		fmt.Println("Cache successfully set!")
	} else {
		fmt.Println("Cache successfully hit!")
		count, err = strconv.ParseInt(countStr, 10, 64)

		if err != nil {
			// Log error
			return 0, nil, err
		}
	}

	blogs, err := db.GetBlogsFromFollowees(ctx, params)
	if err != nil {
		return 0, nil, err
	}

	return count, blogs, nil
}
