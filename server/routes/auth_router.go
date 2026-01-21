package routes

import (
	"crud-app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	user := router.Group("/auth")
	{
		user.POST("/signup", controllers.UserSignup)
		user.POST("/login", controllers.UserLogin)
	}
}
