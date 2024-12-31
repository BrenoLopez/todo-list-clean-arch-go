package routes

import (
	"todolist/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type TodoRoutes struct {
	TodoController controllers.ITodoController
}

type ITodoRoutes interface {
	InitTodoRoutes(httpServer *gin.Engine, todoRoutes *TodoRoutes)
}

func InitRoutes(httpServer *gin.Engine, todoRoutes *TodoRoutes) {
	todoGroup := httpServer.Group("/todos")

	todoGroup.POST("/", todoRoutes.TodoController.Create)
	todoGroup.DELETE("/:id", todoRoutes.TodoController.Delete)
}
