package routes

import (
	controller "finance_backend/controllers/Auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", controller.Login)
}
