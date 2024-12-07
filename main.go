
package main

import (
	"go-todo/config"
	"go-todo/controller"

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

	r.GET("/todos", controller.GetTodos)
	r.POST("/todos", controller.CreateTodo)
	r.GET("/todos/:id", controller.GetTodoById)
	r.PUT("/todos/:id", controller.UpdateTodoById)
	r.DELETE("/todos/:id", controller.DeleteTodoById)

	r.Run()
}