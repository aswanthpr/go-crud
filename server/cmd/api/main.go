package main

import (
	"crud-app/configs"
	"crud-app/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv()
	configs.ConnectDB()

	port := configs.GetEnv("PORT")
	r := gin.Default()

	routes.UserRouter(r)
	log.Println("server is running in  8080")
	r.Run(":" + port)
}
	// router.GET("/ping", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })