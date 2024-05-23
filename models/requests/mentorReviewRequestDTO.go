package requests

import "github.com/google/uuid"

type GetMentorReviewsRequestDTO struct {
	ID string `json:"id" form: "id" binding: "required"`
}

type CreateMentorReviewRequestDTO struct {
	MentorID    uuid.UUID `json:"mentor_id" form: "mentor_id" binding: "required"`
	Description string    `json:"description" form: "description" binding: "omitempty"`
	IsActive    bool      `json:"is_active" form: "is_active" binding: "required"`
	CreatedBy   string    `json:"created_by" form: "created_by" binding: "required"`
}

type UpdateMentorReviewRequestDTO struct {
	ID          string `json:"id" form: "id" binding: "required"`
	MentorID    string `json:"mentor_id" form: "mentor_id" binding: "omitempty"`
	Description string `json:"description" form: "description" binding: "omitempty"`
	IsActive    bool   `json:"is_active" form: "is_active" binding: "omitempty"`
	UpdatedBy   string `json:"updated_by" form: "updated_by" binding: "omitempty"`
}