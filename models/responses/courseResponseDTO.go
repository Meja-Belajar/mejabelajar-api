package responses

import (
	"time"
)

type CourseResponseDTO struct {
	CourseID        string  `json:"course_id"`
	Name            string    `json:"name"`
	Detail          string    `json:"detail"`
	Rating          float64   `json:"rating"`
	HourlyRate      float64   `json:"hourly_rate"`
	CourseStartTime time.Time `json:"course_start_time"`
	CourseEndTime   time.Time `json:"course_end_time"`
}
