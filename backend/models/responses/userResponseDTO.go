package responses

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	UserName       string    `json:"username"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Description    string    `json:"description"`
	ProfilePicture string    `json:"profile_picture"`
	BOD            time.Time `json:"bod"`
	IsAvailable    bool      `json:"is_available"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
