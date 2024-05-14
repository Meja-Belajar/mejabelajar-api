package responses

import (
	"time"

	"github.com/google/uuid"
)

type BookingResponseDTO struct {
	ID     uuid.UUID   `json:"id"`
	User   userBooking `json:"user"`
	Mentor userBooking `json:"mentor"`
	Course struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `josn:"name"`
		Detail string    `json:"detail"`
	} `json:"course"`
	Invoice struct {
		ID             uuid.UUID `json:"id"`
		Payment_method string    `json:"payment_method"`
		Payment_name   string    `json:"payment_name"`
		Payment_status string    `json:"payment_status"`
		Payment_amount float64   `json:"payment_amount"`
		Payment_fee    float64   `json:"payment_fee"`
		Payment_total  float64   `json:"payment_total"`
	} `json:"invoice"`
	BookingDate     time.Time `json:"date"`
	BookingLocation string    `json:"location"`
}

type userBooking struct {
	ID   uuid.UUID `json:"id"`
	Name string    `josn:"name"`
}
