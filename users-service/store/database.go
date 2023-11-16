package store

import "github.com/kylehipz/simpleng-blog-app/blog-post-service/types"

type Database interface {
	FindUsers() ([]*types.User, error)
}
