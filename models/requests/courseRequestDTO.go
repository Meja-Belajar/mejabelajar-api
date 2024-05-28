package requests

import "time"

type GetCourseRequestDTO struct {
	CourseID string `json:"id" form:"id" binding:"required"` // PK
}

type AddCourseRequestDTO struct {
	Name            string    `json:"name" form:"name" binding:"required"`
	Detail          string    `json:"detail" form:"detail" binding:"required"`
	Rating          float64   `json:"rating" form:"rating" binding:"required"`
	HourlyRate      float64   `json:"hourlyrate" form:"hourlyrate" binding:"required"`
	CourseStartTime time.Time `json:"coursestarttime" form:"coursestarttime" binding:"required"`
	CourseEndTime   time.Time `json:"courseendtime" form:"courseendtime" binding:"required"`
	IsActive        bool      `json:"isactive" form:"isactive" binding:"required"`
	CreatedBy       string    `json:"createdby" form:"createdby" binding:"required"`
}

type UpdateCourseRequestDTO struct {
	CourseID        string `json:"id" form:"id" binding:"required"`
	Name            string `json:"name" form:"name" binding:"omitempty"`
	Detail          string `json:"detail" form:"detail" binding:"omitempty"`
	Rating          float64 `json:"rating" form:"rating" binding:"omitempty"`
	HourlyRate      float64 `json:"hourlyrate" form:"hourlyrate" binding:"omitempty"`
	CourseStartTime time.Time `json:"coursestarttime" form:"coursestarttime" binding:"omitempty"`
	CourseEndTime   time.Time `json:"courseendtime" form:"courseendtime" binding:"omitempty"`
	IsActive        bool   `json:"isactive" form:"isactive" binding:"omitempty"`
	UpdatedBy       string `json:"updatedby" form:"updatedby" binding:"required"`
}
