package main

import (
	"fmt"
	"os"
	"todolist/pkg/application"
	"todolist/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file load")
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("app port not set")
	}

	httpServer := gin.Default()
	gin.SetMode(gin.DebugMode)
	httpServer.SetTrustedProxies(nil)

	databaseConnection := database.Start()

	defer func() {
		connection, _ := databaseConnection.DB()
		connection.Close()
	}()

	application.New(httpServer, databaseConnection).Start(fmt.Sprintf(":%s", port))

}
