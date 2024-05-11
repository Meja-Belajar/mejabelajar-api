package outputs

import "github.com/meja_belajar/models/responses"

type BookingByUserIdOutput struct {
	BaseOutput
	Data []responses.BookingResponseDTO `json:"data"`
}

type BookingByBookIdOutput struct {
	BaseOutput
	Data responses.BookingResponseDTO `json:"data"`
}
