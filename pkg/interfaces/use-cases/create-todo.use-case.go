package interfaces

import (
	"todolist/pkg/models"
)

type ICreateTodoUseCase interface {
	Create(data *models.TodoModel) (*models.TodoModel, error)
}
