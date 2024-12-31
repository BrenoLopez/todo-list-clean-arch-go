package request

import "time"

type TodoRequestDto struct {
	Description string    `json:"description,omitempty" binding:"required"`
	DateToDone  time.Time `json:"dateToDone,omitempty" binding:"required"`
}
