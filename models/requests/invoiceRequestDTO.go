package requests

type UpdateInvoiceStatusRequestDTO struct {
	ID            string `json:"id"`
	PaymentStatus string `json:"payment_status"`
}
