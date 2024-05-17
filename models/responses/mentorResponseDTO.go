package responses

import (
	"time"

	"github.com/google/uuid"
)

type MentorResponseDTO struct {
	ID                 uuid.UUID           `json:"mentor_id"`
	UserName           string              `json:"username"`
	University         string              `json:"university"`
	Email              string              `json:"email"`
	Phone              string              `json:"phone"`
	Description        string              `json:"description"`
	ProfilePicture     string              `json:"profile_picture"`
	BOD                time.Time           `json:"bod"`
	Revenue            float64             `json:"revenue"`
	Rating             float64             `json:"rating"`
	TotalTeachingHours float64             `json:"total_teaching_hours"`
	TeachingFrequency  float64             `json:"teaching_frequency"`
	Courses            []CourseResponseDTO `json:"courses"`
	Reviews            []ReviewResponseDTO `json:"reviews"`
}
