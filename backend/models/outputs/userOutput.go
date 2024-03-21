package outputs

import "github.com/meja_belajar/models/responses" // Import the missing package

type RegisterUserOutput struct {
	BaseOutput
	Data responses.UserResponseDTO `json:"data"`
}
