package usecases

import (
	"simpleng-blog-app/common/pkg/database"
	"simpleng-blog-app/common/pkg/models"
)

func RegisterUser(username string) (*models.User, error) {
	newUser := models.User{UserName: username}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
