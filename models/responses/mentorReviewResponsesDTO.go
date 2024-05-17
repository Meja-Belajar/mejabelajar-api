package responses

import (
	"time"

	"github.com/google/uuid"
)

type MentorReviewsResponseDTO struct {
	ID          uuid.UUID    `json:"id"`
	MentorID    uuid.UUID    `json:"mentor_id"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}