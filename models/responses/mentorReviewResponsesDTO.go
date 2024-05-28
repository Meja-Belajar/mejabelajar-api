package responses

import (
	"github.com/google/uuid"
)

type MentorReviewsResponseDTO struct {
	ID          uuid.UUID   `json:"id"`
	Description string    	`json:"description"`
}