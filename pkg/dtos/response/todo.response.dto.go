package response

import (
	"time"

	"github.com/google/uuid"
)

type TodoReponseDto struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	DateToDone  time.Time `json:"dateToDone,omitempty"`
}
