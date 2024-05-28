package helpers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/utils"

	// "github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
)

func FindInvoiceByUserID(userID string, ctx context.Context) (int, interface{}) {
	db := configs.GetDB()
	rows, err := db.Table("invoices i").
		WithContext(ctx).
		Select("i.id, i.payment_method, i.payment_name, i.payment_status, i.payment_amount, i.payment_fee, i.payment_total").
		Joins("JOIN bookings b ON b.invoice_id = i.id").
		Joins("JOIN users u ON b.user_id = u.id").
		Where("u.id = ?", userID).
		Rows()

	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return utils.HandleNotFound("Invoice")
	}

	var invoices []responses.InvoiceResponseDTO
	var invoice responses.InvoiceResponseDTO
	for {
		err := rows.Scan(&invoice.ID, &invoice.Payment_method, &invoice.Payment_name,
			&invoice.Payment_status, &invoice.Payment_amount, &invoice.Payment_fee,
			&invoice.Payment_total)

		if err != nil {
			return utils.HandleInternalServerError(err)
		}

		invoices = append(invoices, invoice)
		if !rows.Next() {
			break
		}
	}

	output := outputs.InvoiceByUserIDOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: List of invoices has been retrieved",
		},
		Data: invoices,
	}
	return 200, output
}

func FindAllInvoice(ctx context.Context) (int, interface{}) {
	db := configs.GetDB()
	rows, err := db.Table("invoices i").
		WithContext(ctx).
		Select("i.id, i.payment_method, i.payment_name, i.payment_status, i.payment_amount, i.payment_fee, i.payment_total").
		Rows()

	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return utils.HandleNotFound("Invoice")
	}

	var invoices []responses.InvoiceResponseDTO
	var invoice responses.InvoiceResponseDTO
	for {
		err := rows.Scan(&invoice.ID, &invoice.Payment_method, &invoice.Payment_name,
			&invoice.Payment_status, &invoice.Payment_amount, &invoice.Payment_fee,
			&invoice.Payment_total)

		if err != nil {
			return utils.HandleInternalServerError(err)
		}

		invoices = append(invoices, invoice)
		if !rows.Next() {
			break
		}
	}

	output := outputs.InvoiceByUserIDOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: List of invoices has been retrieved",
		},
		Data: invoices,
	}
	return 200, output
}

func FindInvoiceByInvoiceID(invoiceID string) (int, interface{}) {
	db := configs.DB
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var invoice responses.InvoiceResponseDTO
	rows, err := db.Table("invoices i").
		WithContext(ctx).
		Select("i.id, i.payment_method, i.payment_name, i.payment_status, i.payment_amount, i.payment_fee, i.payment_total").
		Where("i.id = ?", invoiceID).
		Rows()

	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return utils.HandleNotFound("Invoice")
	}

	err = rows.Scan(&invoice.ID, &invoice.Payment_method, &invoice.Payment_name, &invoice.Payment_status, &invoice.Payment_amount, &invoice.Payment_fee, &invoice.Payment_total)

	if err != nil {
		return utils.HandleInternalServerError(err)
	}

	output := outputs.InvoiceByInvoiceIDOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Invoice with ID has been retrieved",
		},
		Data: invoice,
	}
	return 200, output
}

func UpdateInvoiceStatus(UpdateInvoiceStatusRequestDTO requests.UpdateInvoiceStatusRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var invoice database.Invoices

	err := db.First(&invoice, "id = ?", utils.StringToUUID(UpdateInvoiceStatusRequestDTO.ID)).Error
	if err != nil {
		return utils.HandleInternalServerError(err)
	}

	if invoice.ID == uuid.Nil {
		return utils.HandleNotFound("Invoice")
	}

	invoice.PaymentStatus = UpdateInvoiceStatusRequestDTO.PaymentStatus

	err = db.Save(&invoice).Error
	if err != nil {
		return utils.HandleInternalServerError(err)
	}

	output := outputs.UpdateInvoiceOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.InvoiceResponseDTO{
		ID:             invoice.ID,
		Payment_method: invoice.PaymentMethod,
		Payment_name:   invoice.PaymentName,
		Payment_status: invoice.PaymentStatus,
		Payment_amount: invoice.PaymentAmount,
		Payment_fee:    invoice.PaymentFee,
		Payment_total:  invoice.PaymentTotal,
	}
	return 200, output
}
