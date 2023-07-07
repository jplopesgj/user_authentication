package routes

import (
	userRoutes "finance_backend/routes/User"

	"github.com/gin-gonic/gin"
)

type Route struct {
	userRoute userRoutes.UserRoute
}

func (ur *Route) HandleRequest() {
	router := gin.New()
	api := router.Group("/api")
	ur.userRoute.UserRoutes(api)
	router.Run(":8080")
}
