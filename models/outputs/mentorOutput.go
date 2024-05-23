package outputs

import "github.com/meja_belajar/models/responses"

type RegisterMentor struct {
	BaseOutput
	Data responses.MentorResponseDTO `json:"data"`
}

type GetMentorByIDOutput struct {
	BaseOutput
	Data responses.MentorResponseDTO `json:"data"`
}

type GetAllMentorOutput struct {
	BaseOutput
	Data []responses.MentorResponseDTO `json:"data"`
}


type GetPopularMentorOutput struct {
	BaseOutput
	Data []responses.MentorResponseDTO `json:"data"`
}