package routes

import (
	userController "user_auth/controllers/User"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController userController.UserController
}

func (ur *UserRoute) UserRoutes(api *gin.RouterGroup) {
	api.POST("/user/register", ur.userController.RegisterUser)
	api.GET("/user/:id", ur.userController.GetSpecificUser)
	api.PATCH("/user/:id", ur.userController.UpdateUser)
	api.DELETE("/user/:id", ur.userController.DeleteUser)
}
