package routes

import (
	"crud-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("/signup", controllers.UserSignup)
	}
}
