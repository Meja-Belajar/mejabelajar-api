package responses

import "github.com/google/uuid"

type invoiceResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	Payment_method string    `json:"payment_method"`
	Payment_name   string    `json:"payment_name"`
	Payment_status string    `json:"payment_status"`
	Payment_amount float64   `json:"payment_amount"`
	Payment_fee    float64   `json:"payment_fee"`
	Payment_total  float64   `json:"payment_total"`
}