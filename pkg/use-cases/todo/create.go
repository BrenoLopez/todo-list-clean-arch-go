package todo

import (
	interfaces "todolist/pkg/interfaces/repositories"
	"todolist/pkg/models"
)

type createTodo struct {
	todoRepository interfaces.ITodoRepository
}

func NewCreateTodoUseCase(todoRepository interfaces.ITodoRepository) *createTodo {
	return &createTodo{todoRepository: todoRepository}
}

func (useCase *createTodo) Create(data *models.TodoModel) (*models.TodoModel, error) {
	result, err := useCase.todoRepository.Create(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
