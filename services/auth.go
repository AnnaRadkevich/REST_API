package services

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"main/database"
	"main/models"
	"main/utils"
)

func SighUp(userInput models.UserRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	var user = models.User{
		ID:       uuid.New().String(),
		Email:    userInput.Email,
		Password: string(password),
	}
	database.DB.Create(&user)
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(userInput models.UserRequest) (string, error) {
	var user models.User
	result := database.DB.First(&user, "email = ?", userInput.Email)
	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return "", err
	}
	return token, nil
}
