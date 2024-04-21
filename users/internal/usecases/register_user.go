package usecases

import (
	"simpleng-blog-app/users/internal/database"
	"simpleng-blog-app/users/internal/models"
)

func RegisterUser(username string) (*models.User, error) {
	newUser := models.User{UserName: username}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
