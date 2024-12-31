package database

import (
	"fmt"
	"os"
	"todolist/pkg/database/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Start() (databaseConnection *gorm.DB) {

	stringConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	databaseConnection, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: false,
		}})

	if (err) != nil {
		panic(err)
	}

	config, _ := databaseConnection.DB()

	config.SetMaxOpenConns(100)
	config.SetMaxIdleConns(10)

	initMigration(databaseConnection)

	fmt.Println("Database connected")
	return databaseConnection
}

func initMigration(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.Table("todos").AutoMigrate(&entities.TodoEntity{})
}
