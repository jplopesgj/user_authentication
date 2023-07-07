package services

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetSpecificUser(c *gin.Context)
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
