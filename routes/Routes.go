package routes

import (
	authRoutes "user_auth/routes/Auth"
	userRoutes "user_auth/routes/User"

	"github.com/gin-gonic/gin"
)

type Route struct {
	userRoute userRoutes.UserRoute
	authRoute authRoutes.AuthRoute
}

func (ur *Route) HandleRequest() {
	router := gin.New()
	api := router.Group("/api")
	ur.userRoute.UserRoutes(api)
	ur.authRoute.AuthRoutes(api)

	router.Run(":8080")

	//.Use(middleware.Auth()) To check if the token is accepted, use this middleware.
}
