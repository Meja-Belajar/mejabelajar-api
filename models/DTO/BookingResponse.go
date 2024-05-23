package DTO

import (
	"github.com/meja_belajar/models/responses"
)

func ToBookingResponse(booking responses.BookingData) responses.BookingResponseDTO {
	bookingResponse := responses.BookingResponseDTO{
		ID:              booking.BookingID,
		BookingDate:     booking.BookingDate,
		BookingLocation: booking.Location,
	}

	bookingResponse.User = responses.UserBooking{
		ID:   booking.UserID,
		Name: booking.UserName,
	}

	bookingResponse.Mentor = responses.UserBooking{
		ID:   booking.MentorID,
		Name: booking.MentorName,
	}

	bookingResponse.Course.ID = booking.CourseID
	bookingResponse.Course.Name = booking.CourseName
	bookingResponse.Course.Detail = booking.CourseDetail

	bookingResponse.Invoice.ID = booking.InvoiceID
	bookingResponse.Invoice.Payment_method = booking.InvoicePaymentMethod
	bookingResponse.Invoice.Payment_name = booking.InvoicePaymentName
	bookingResponse.Invoice.Payment_status = booking.InvoicePaymentStatus
	bookingResponse.Invoice.Payment_amount = booking.InvoicePaymentAmount
	bookingResponse.Invoice.Payment_fee = booking.InvoicePaymentFee
	bookingResponse.Invoice.Payment_total = booking.InvoicePaymentTotal
	return bookingResponse
}

func ToBookingResponses(listBooking []responses.BookingData) []responses.BookingResponseDTO {
	var bookingResponses []responses.BookingResponseDTO = nil

	for _, booking := range listBooking {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}
	return bookingResponses
}
