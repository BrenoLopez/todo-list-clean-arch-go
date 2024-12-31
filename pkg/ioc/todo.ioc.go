package ioc

import (
	"sync"
	"todolist/pkg/controllers"
	"todolist/pkg/repositories"
	"todolist/pkg/use-cases/todo"

	"gorm.io/gorm"
)

var once sync.Once

func NewTodoModule(dataSource *gorm.DB) controllers.ITodoController {
	todoRepository := repositories.NewTodoRepository(dataSource)
	controller := controllers.NewTodoController(nil, nil)

	once.Do(func() {
		controller = controllers.NewTodoController(todo.NewCreateTodoUseCase(todoRepository), todo.NewDeleteTodoUseCase(todoRepository))
	})

	return controller
}
