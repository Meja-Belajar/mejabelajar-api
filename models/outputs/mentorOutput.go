package outputs

import "github.com/meja_belajar/models/responses"

type MentorOutput struct {
	BaseOutput
	Data responses.MentorResponseDTO `json:"data"`
}
