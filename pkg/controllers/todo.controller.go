package controllers

import (
	"net/http"
	"todolist/pkg/dtos/request"
	"todolist/pkg/dtos/response"
	"todolist/pkg/interfaces"
	"todolist/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITodoController interface {
	Create(httpContext *gin.Context)
	Delete(httpContext *gin.Context)
}

type todoController struct {
	createTodoUseCase interfaces.ICreateTodoUseCase
	deleteTodoUseCase interfaces.IDeleteTodoUseCase
}

func NewTodoController(createTodoUseCase interfaces.ICreateTodoUseCase, deleteTodoUseCase interfaces.IDeleteTodoUseCase) *todoController {
	return &todoController{createTodoUseCase: createTodoUseCase, deleteTodoUseCase: deleteTodoUseCase}
}

func (todoController *todoController) Create(httpContext *gin.Context) {
	var todoRequestDto request.TodoRequestDto
	if err := httpContext.ShouldBindJSON(&todoRequestDto); err != nil {
		panic(err)
	}

	result, _ := todoController.createTodoUseCase.Create(models.NewTodo(todoRequestDto.Description, todoRequestDto.DateToDone))

	httpContext.JSON(http.StatusOK, response.TodoReponseDto{Id: result.Id, Description: result.Description, DateToDone: result.DateToDone})
}

func (todoController *todoController) Delete(httpContext *gin.Context) {
	id := httpContext.Param("id")

	if err := uuid.Validate(id); err != nil {
		httpContext.JSON(http.StatusBadGateway, gin.H{"message": "Invalid UUID!"})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		panic("Error on parse uuid")
	}
	todoController.deleteTodoUseCase.Delete(uuid)

	httpContext.Status(http.StatusOK)
}
