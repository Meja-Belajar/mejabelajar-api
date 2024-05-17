package outputs

import (
	"github.com/meja_belajar/models/responses"
)

type GetAllMentorReviewByMentorID struct {
	BaseOutput
	Data []responses.MentorReviewsResponseDTO `json:"data"`
}

type GetMentorReviewsOutput struct {
	BaseOutput
	Data responses.MentorReviewsResponseDTO `json:"data"`
}

type CreateMentorReviewRequestDTO struct {
	BaseOutput
	Data responses.MentorReviewsResponseDTO `json:"data"`
}

type UpdateMentorReviewRequestDTO struct {
	BaseOutput
	Data responses.MentorReviewsResponseDTO `json:"data"`
}