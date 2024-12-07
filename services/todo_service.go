package service

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

func GetTaskById(id int) (models.Todo, error) {
	var todo models.Todo
	err := config.DB.First(&todo, id).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func UpdateTaskById(id int, input models.Todo) (models.Todo, error) {
	todo, err := GetTaskById(id)
	if err != nil {
		return todo, err
	}

	err = config.DB.Model(&todo).Updates(input).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func DeleteTaskById(id int) error {
	todo, err := GetTaskById(id)
	if err != nil {
		return err
	}

	err = config.DB.Delete(&todo).Error
	if err != nil {
		return err
	}

	return nil
}