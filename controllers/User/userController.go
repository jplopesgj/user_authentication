package controllers

import (
	services "finance_backend/service/User"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) GetSpecificUser(c *gin.Context) {
	uc.GetSpecificUser(c)
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	uc.RegisterUser(c)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	uc.UpdateUser(c)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	uc.DeleteUser(c)
}
