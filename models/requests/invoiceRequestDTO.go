package requests

type invoiceRequestDTO struct {
	PaymentMethod 	string `json : "payment_method"`
	PaymentName 	string `json : "payment_name"`
	PaymentStatus 	string `json : "payment_status"`
	PaymentAmount 	float64 `json : "payment_amount"`
}
