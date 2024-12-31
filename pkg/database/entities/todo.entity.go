package entities

import (
	"time"

	"github.com/google/uuid"
)

type TodoEntity struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	DateToDone  time.Time `gorm:"type:timestamptz;not null"`
}

type Tabler interface {
	TableName() string
}

func (TodoEntity) TableName() string {
	return "todos"
}
