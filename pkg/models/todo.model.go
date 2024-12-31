package models

import (
	"time"

	"github.com/google/uuid"
)

type TodoModel struct {
	Id          uuid.UUID
	Description string
	DateToDone  time.Time
}

func NewTodo(description string, dateToDone time.Time) *TodoModel {
	return &TodoModel{
		Id: uuid.New(), Description: description, DateToDone: dateToDone,
	}
}
