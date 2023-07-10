package controllers

import (
	service "user_auth/service/User"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserServiceImpl
}

func (uc *UserController) GetSpecificUser(c *gin.Context) {
	uc.userService.GetSpecificUser(c)
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	uc.userService.RegisterUser(c)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	uc.userService.UpdateUser(c)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	uc.userService.DeleteUser(c)
}
