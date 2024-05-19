package outputs

import "github.com/meja_belajar/models/responses"

type GetMentorByIDOutput struct {
	BaseOutput
	Data responses.MentorResponseDTO `json:"data"`
}
