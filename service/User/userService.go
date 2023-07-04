package service

import (
	"finance_backend/database"
	models "finance_backend/models/User"
	"time"
)

func CreateUser(userInput models.CreateUserInput) models.User {
	user := models.User{
		Name:      userInput.Name,
		Lastname:  userInput.Lastname,
		Password:  userInput.Password,
		CreatedAt: time.Now(),
	}

	database.DB.Create(&user)

	return user
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
