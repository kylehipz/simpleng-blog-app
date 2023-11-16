package services

import (
	"log"

	"github.com/kylehipz/simpleng-blog-app/blog-post-service/store"
	"github.com/kylehipz/simpleng-blog-app/blog-post-service/types"
)

type UsersService struct {
	DB store.Database
}

func (s *UsersService) GetUsers() ([]*types.User, error) {
	users, err := s.DB.FindUsers()
	if err != nil {
		log.Fatalln(err)
	}

	return users, nil
}
