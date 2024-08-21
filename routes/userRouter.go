package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tokha04/jwt-authentication/controllers"
	"github.com/tokha04/jwt-authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
