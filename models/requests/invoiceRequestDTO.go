package requests

type UpdateInvoiceStatusRequestDTO struct {
	ID             string `json : "id"`
	PayementStatus string `json : "payment_status"`
}
