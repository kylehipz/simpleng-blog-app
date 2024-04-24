package usecases

import (
	"context"

	"simpleng-blog-app/common/pkg/database"
)

func RegisterUser(username string) (*database.User, error) {
	db := database.DB

	newUser, err := db.CreateUser(context.Background(), username)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
