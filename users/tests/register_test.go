package main

import (
	"os"
	"testing"

	"simpleng-blog-app/users/internal/database"
	"simpleng-blog-app/users/internal/models"
	"simpleng-blog-app/users/internal/usecases"
)

func TestMain(m *testing.M) {
	database.Connect()

	db := database.DB
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	code := m.Run()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	username := "kylehipz"

	newUser, err := usecases.RegisterUser(username)
	if err != nil {
		t.Errorf("RegisterUser should be able to register user %s\n", username)
	}

	// Fetch record
	db := database.DB
	user := models.User{}
	db.First(&user, "user_name = ?", username)

	if user.UserName != username && user.UserName != newUser.UserName &&
		newUser.UserName != username {
		t.Errorf("RegisterUser should be able to register %s\n", username)
	}
}

func TestRegisterUniqueError(t *testing.T) {
	username := "kylehipolito"

	_, err := usecases.RegisterUser(username)
	if err != nil {
		t.Errorf("Register user failed")
	}

	_, err = usecases.RegisterUser(username)

	if err == nil {
		t.Errorf("RegisterUser should not be able to register user %s again", username)
	}
}
