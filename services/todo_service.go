package services

import (
	"go-todo/config"
	"go-todo/models"
)

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := config.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodo(input models.Todo) (models.Todo, error) {
	err := config.DB.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}