package responses

import (
	"github.com/google/uuid"
)

type SearchResponseDTO struct {
	MentorId       uuid.UUID `json:"mentor_id"`
	Username       string    `json:"username"`
	University     string    `json:"university"`
	ProfilePicture string    `json:"profile_picture"`
	Rating         float64   `json:"rating"`
}
