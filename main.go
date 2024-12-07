package main

import (
	"go-todo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}