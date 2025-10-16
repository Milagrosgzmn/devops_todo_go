package repository

import "github.com/Milagrosgzmn/devops_todo_go.git/internal/models"

type IRepository interface {
	Get(id int) (models.TodoItem, error)
	GetAll() (	[]models.TodoItem, error)
	Create(item models.TodoItem)(models.TodoItem, error)
	Update(item models.TodoItem) error
	Delete(id string) error

}