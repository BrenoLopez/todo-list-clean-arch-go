package interfaces

import (
	"todolist/pkg/models"

	"github.com/google/uuid"
)

type ITodoRepository interface {
	Create(input *models.TodoModel) (*models.TodoModel, error)
	Delete(id uuid.UUID) error
}
