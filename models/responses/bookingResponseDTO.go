package responses

import (
	"time"

	"github.com/google/uuid"
)

type BookingResponseDTO struct {
	ID     uuid.UUID   `json:"id"`
	User   UserBooking `json:"user"`
	Mentor UserBooking `json:"mentor"`
	Course struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
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

type UserBooking struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type BookingData struct {
	BookingID   uuid.UUID
	BookingDate time.Time
	Location    string

	// User
	UserID   uuid.UUID `gorm:"user_id"`
	UserName string

	// Mentor
	MentorID   uuid.UUID
	MentorName string

	// Course
	CourseID     uuid.UUID
	CourseName   string
	CourseDetail string

	// Invoice
	InvoiceID            uuid.UUID
	InvoicePaymentMethod string
	InvoicePaymentName   string
	InvoicePaymentStatus string
	InvoicePaymentAmount float64
	InvoicePaymentFee    float64
	InvoicePaymentTotal  float64
}
