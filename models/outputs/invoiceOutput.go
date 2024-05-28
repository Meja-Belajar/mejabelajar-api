package outputs

import "github.com/meja_belajar/models/responses"

type InvoiceByUserIDOutput struct{
	BaseOutput
	Data []responses.InvoiceResponseDTO `json:"data"`
}

type InvoiceByInvoiceIDOutput struct{
	BaseOutput
	Data responses.InvoiceResponseDTO `json:"data"`
}

type UpdateInvoiceOutput struct{
	BaseOutput
	Data responses.InvoiceResponseDTO `json:"data"`
}