package interfaces

import "github.com/google/uuid"

type IDeleteTodoUseCase interface {
	Delete(id uuid.UUID) error
}
