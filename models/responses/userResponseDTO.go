package responses

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	UserName       string    `json:"username"`
	University     string    `json:"university"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Description    string    `json:"description"`
	ProfilePicture string    `json:"profile_picture"`
	BOD            time.Time `json:"bod"`
	IsActive       bool      `json:"is_active"`
	IsMentor       bool      `json:"is_mentor"`
}
