package outputs

import "github.com/meja_belajar/models/responses" // Import the missing package

type RegisterUserOutput struct 
	BaseOutput
	Data responses.UserResponseDTO `json:"data"`
}

type LoginUserOutput struct {
	BaseOutput
	Data responses.UserResponseDTO `json:"data"`
}

type GetUserByIDOutput struct{
	BaseOutput
	Data responses.UserResponseDTO `json:"data"`
}
