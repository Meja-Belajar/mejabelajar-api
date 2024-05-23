package responses

import (
	"time"
)

type MentorResponseDTO struct {
	MentorID           string              `json:"mentor_id"`
	Username           string              `json:"username"`
	University         string              `json:"university"`
	Email              string              `json:"email"`
	Phone              string              `json:"phone"`
	Description        string              `json:"description"`
	ProfilePicture     string              `json:"profile_picture"`
	BOD                time.Time           `json:"bod"`
	Revenue            float64             `json:"revenue"`
	Rating             float64             `json:"rating"`
	TotalTeachingHours int             `json:"total_teaching_hours"`
	TeachingFrequency  int             `json:"teaching_frequency"`
	Courses            []CourseResponseDTO `json:"courses"`
	Reviews            []ReviewResponseDTO `json:"reviews"`
}
