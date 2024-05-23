package outputs

import "github.com/meja_belajar/models/responses"

// By ID
type GetCourseOutput struct {
	BaseOutput
	Data responses.CourseResponseDTO `json:"data"`
}

type AddCourseOutput struct {
	BaseOutput
	Data responses.CourseResponseDTO `json:"data"`
}

type UpdateCourseOutput struct {
	BaseOutput
	Data responses.CourseResponseDTO `json:"data"`
}
