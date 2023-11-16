package store

import "github.com/kylehipz/simpleng-blog-app/users-service/types"

type Database interface {
	FindUsers() ([]*types.User, error)
}
