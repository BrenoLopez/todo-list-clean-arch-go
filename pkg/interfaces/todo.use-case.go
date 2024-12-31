package interfaces

import (
	"todolist/pkg/models"

	"github.com/google/uuid"
)

type ITodoUseCase interface {
	Create(data *models.TodoModel) (*models.TodoModel, error)
	Delete(id uuid.UUID) error
}
