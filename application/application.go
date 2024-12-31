package application

import (
	"todolist/pkg/ioc"
	"todolist/pkg/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IApplication interface {
	Start(port string)
}

type application struct {
	ginEngine          *gin.Engine
	databaseConnection *gorm.DB
}

func New(ginEngine *gin.Engine, databaseConnection *gorm.DB) IApplication {
	return &application{ginEngine: ginEngine, databaseConnection: databaseConnection}
}

func (server *application) Start(port string) {

	routes.InitRoutes(server.ginEngine, &routes.TodoRoutes{TodoController: ioc.NewTodoModule(server.databaseConnection)})
	server.ginEngine.Run(port)
}
