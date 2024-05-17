package responses

import (
	"time"

	"github.com/google/uuid"
)

type CourseResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	IsActive  bool      `json:"isactive"`
	CreatedBy string    `json:"createdby"`
	UpdatedBy string    `json:"updatedby"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
