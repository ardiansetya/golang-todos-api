package controller

import (
	"go-todo/services"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Get all todos Successfully",
		"data": todos,
	})
}