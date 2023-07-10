package routes

import (
	controller "user_auth/controllers/Auth"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController controller.AuthController
}

func (ar *AuthRoute) AuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", ar.authController.Login)
}
