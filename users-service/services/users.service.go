package services

import (
	"github.com/kylehipz/simpleng-blog-app/users-service/store"
	"github.com/kylehipz/simpleng-blog-app/users-service/types"
)

type UsersService struct {
	DB store.Database
}

func (s *UsersService) GetUsers() ([]*types.User, error) {
	users, err := s.DB.FindUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UsersService) Follow(follow *types.Follow) (*types.Follow, error) {
	inserted, err := s.DB.InsertFollow(follow)
	if err != nil {
		return nil, err
	}

	return inserted, nil
}
