package repositories

import (
	"todolist/pkg/database/entities"
	interfaces "todolist/pkg/interfaces/repositories"
	"todolist/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type todoRepository struct {
	dataSource *gorm.DB
}

func NewTodoRepository(dataSource *gorm.DB) interfaces.ITodoRepository {
	return &todoRepository{dataSource: dataSource}
}

func (todoRepository *todoRepository) Create(data *models.TodoModel) (*models.TodoModel, error) {
	todo := &entities.TodoEntity{Description: data.Description, DateToDone: data.DateToDone}
	result := todoRepository.dataSource.Create(todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.TodoModel{Id: todo.Id, Description: todo.Description, DateToDone: todo.DateToDone}, nil
}

func (todoRepository *todoRepository) Delete(id uuid.UUID) error {
	todoRepository.dataSource.Delete(&entities.TodoEntity{}, id)
	return nil
}
