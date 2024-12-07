package controller

import (
	"go-todo/models"
	"go-todo/services"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := service.GetAllTodos()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todos have been fetched successfully",
		"data":    todos,
	})
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	todo, err := service.CreateTodo(input)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"success": true,
		"message": "Todo has been created successfully",
		"data":    todo,
	})
}

func GetTodoById(c *gin.Context) {
	todo_id := c.Param("id")
	id, err := strconv.Atoi(todo_id)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	todo, err := service.GetTaskById(id)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo has been fetched successfully",
		"data":    todo,
	})
}

func UpdateTodoById(c *gin.Context) {
	var input models.Todo

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	todo_id := c.Param("id")
	id, err := strconv.Atoi(todo_id)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	todo, err := service.UpdateTaskById(id, input)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo has been updated successfully",
		"data":    todo,
	})

}

func DeleteTodoById(c *gin.Context) {
	todo_id := c.Param("id")
	id, err := strconv.Atoi(todo_id)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	err = service.DeleteTaskById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo has been deleted successfully",
	})
}