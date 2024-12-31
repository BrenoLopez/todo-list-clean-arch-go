package todo

import (
	"todolist/pkg/interfaces"

	"github.com/google/uuid"
)

type deleteTodo struct {
	todoRepository interfaces.ITodoRepository
}

func NewDeleteTodoUseCase(todoRepository interfaces.ITodoRepository) *deleteTodo {
	return &deleteTodo{todoRepository: todoRepository}
}

func (useCase *deleteTodo) Delete(id uuid.UUID) {
	err := useCase.todoRepository.Delete(id)

	if err != nil {
		panic(err)
	}
}
