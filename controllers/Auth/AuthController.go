package controller

import (
	"net/http"
	database "user_auth/database/config"
	models "user_auth/models/User"
	services "user_auth/service/Auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authServiceImpl services.AuthServiceImpl
	user            models.User
	request         TokenRequest
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *AuthController) Login(context *gin.Context) {
	if err := context.ShouldBindJSON(&ac.request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.DB.Where("email = ?", ac.request.Email).First(&ac.user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := ac.checkPassword(ac.request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := ac.authServiceImpl.GenerateJWT(ac.user.Email, ac.user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (ac *AuthController) checkPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(ac.user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
