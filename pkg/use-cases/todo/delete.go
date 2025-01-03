package todo

import (
	interfaces "todolist/pkg/interfaces/repositories"

	"github.com/google/uuid"
)

type deleteTodo struct {
	todoRepository interfaces.ITodoRepository
}

func NewDeleteTodoUseCase(todoRepository interfaces.ITodoRepository) *deleteTodo {
	return &deleteTodo{todoRepository: todoRepository}
}

func (useCase *deleteTodo) Delete(id uuid.UUID) error {
	err := useCase.todoRepository.Delete(id)

	if err != nil {
		return err
	}
	return nil
}
