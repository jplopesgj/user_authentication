package routes

import (
	controllers "finance_backend/controllers/User"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.New()
	router.GET("/api/user/:id", controllers.GetSpecificUser)
	router.POST("/api/user", controllers.CreateUser)
	router.PATCH("/api/user/:id", controllers.UpdateUser)
	router.DELETE("/api/user/:id", controllers.DeleteUser)
	router.Run(":8080")
}
