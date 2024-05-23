package helpers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/DTO"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/repositories"
	"gorm.io/gorm"
)

func FindBookingByUserID(userID string, ctx context.Context) (int, interface{}) {
	db := configs.GetDB()
	rows, err := db.Table("bookings b").
		WithContext(ctx).
		Select("b.id, u.id, u.username, TableSub.mentorId, TableSub.mentorName, c.id, c.name, c.detail, i.id, i.payment_method, i.payment_name, i.payment_status, i.payment_amount, i.payment_fee, i.payment_total, b.date, b.location").
		Joins("JOIN users u ON b.user_id = u.id").
		Joins("JOIN courses c ON b.course_id = c.id").
		Joins("JOIN invoices i ON b.invoice_id = i.id").
		Joins("JOIN (SELECT b2.id AS bookingID, m2.id AS mentorID, u2.username AS mentorName FROM bookings b2 JOIN mentors m2 ON b2.mentor_id = m2.id JOIN users u2 ON u2.id = m2.user_id WHERE b2.user_id = ?) AS TableSub ON b.id = TableSub.bookingID", userID).
		Rows()

	if err != nil {
		output := outputs.NotFoundOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}
	defer rows.Close()

	if !rows.Next() {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Booking Not Found",
		}
		return 404, output
	}

	var bookings []responses.BookingResponseDTO
	var booking responses.BookingResponseDTO
	for {
		err := rows.Scan(&booking.ID, &booking.User.ID, &booking.User.Name,
			&booking.Mentor.ID, &booking.Mentor.Name,
			&booking.Course.ID, &booking.Course.Name, &booking.Course.Detail,
			&booking.Invoice.ID, &booking.Invoice.Payment_method, &booking.Invoice.Payment_name, &booking.Invoice.Payment_status, &booking.Invoice.Payment_amount, &booking.Invoice.Payment_fee, &booking.Invoice.Payment_total,
			&booking.BookingDate, &booking.BookingLocation)

		if err != nil {
			output := outputs.NotFoundOutput{
				Code:    500,
				Message: "Internal Server Error",
			}
			return 500, output
		}

		bookings = append(bookings, booking)
		if !rows.Next() {
			break
		}
	}

	output := outputs.BookingByUserIdOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: List of bookings has been retrieved",
		},
		Data: bookings,
	}
	return 200, output
}

func FindBookingByBookingID(bookingID string) (int, interface{}) {
	db := configs.DB
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var booking responses.BookingResponseDTO
	rows, err := db.Table("bookings b").
		WithContext(ctx).
		Select("b.id, u.id, u.username, TableSub.mentorId, TableSub.mentorName, c.id, c.name, c.detail, i.id, i.payment_method, i.payment_name, i.payment_status, i.payment_amount, i.payment_fee, i.payment_total, b.date, b.location").
		Joins("JOIN users u ON b.user_id = u.id").
		Joins("JOIN courses c ON b.course_id = c.id").
		Joins("JOIN invoices i ON b.invoice_id = i.id").
		Joins("JOIN (SELECT b2.id AS bookingID, m2.id AS mentorID, u2.username AS mentorName FROM bookings b2 JOIN mentors m2 ON b2.mentor_id = m2.id JOIN users u2 ON u2.id = m2.user_id WHERE b2.id = ?) AS TableSub ON b.id = TableSub.bookingID", bookingID).
		Rows()

	if err != nil {
		output := outputs.NotFoundOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}
	defer rows.Close()

	if !rows.Next() {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Booking Not Found",
		}
		return 404, output
	}

	err = rows.Scan(&booking.ID, &booking.User.ID, &booking.User.Name,
		&booking.Mentor.ID, &booking.Mentor.Name,
		&booking.Course.ID, &booking.Course.Name, &booking.Course.Detail,
		&booking.Invoice.ID, &booking.Invoice.Payment_method, &booking.Invoice.Payment_name, &booking.Invoice.Payment_status, &booking.Invoice.Payment_amount, &booking.Invoice.Payment_fee, &booking.Invoice.Payment_total,
		&booking.BookingDate, &booking.BookingLocation)

	if err != nil {
		output := outputs.NotFoundOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}
	output := outputs.BookingByBookIdOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Booking with ID has been retrieved",
		},
		Data: booking,
	}
	return 200, output
}

func CreateBooking(ctx context.Context, bookingData requests.NewBookingRequestDTO) (int, interface{}) {
	invoice := database.Invoices{
		PaymentMethod: bookingData.Invoice.Payment_method,
		PaymentName:   bookingData.Invoice.Payment_name,
		PaymentStatus: bookingData.Invoice.Payment_status,
		PaymentAmount: bookingData.Invoice.Payment_amount,
		PaymentFee:    0.007 * bookingData.Invoice.Payment_amount,
	}
	invoice.PaymentTotal = invoice.PaymentAmount + invoice.PaymentFee

	err := configs.DB.Create(&invoice).Error
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Message: "Internal Server Error"}
	}

	parsedDate, err := time.Parse("2006-01-02T15:04:05Z", bookingData.ScheduleTime)
	if err != nil {
		return 400, outputs.BadRequestOutput{Message: "Bad Request: Invalid date format"}
	}

	booking := database.Bookings{
		InvoiceID: invoice.ID,
		Location:  bookingData.ScheduledLocation,
		Date:      parsedDate,
	}

	booking.UserID, err = uuid.Parse(bookingData.UserID)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Message: "error parsing UserID"}
	}

	booking.MentorID, err = uuid.Parse(bookingData.MentorID)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Message: "error parsing Mentor ID"}
	}

	booking.CourseID, err = uuid.Parse(bookingData.CourseID)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Message: "error parsing Course ID"}
	}

	inputtedBooking, err := repositories.CreateBooking(ctx, booking)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Message: "Internal Server Error"}
	}

	code, output := FindBookingByBookingID(inputtedBooking.ID.String())
	if code != 200 {
		return 500, outputs.InternalServerErrorOutput{Message: "Internal Server Error"}
	}

	baseOutput := output.(outputs.BookingByBookIdOutput).BaseOutput
	baseOutput.Code = 201
	baseOutput.Message = "Success: Booking Created"
	err = MakeNotification(ctx, output.(outputs.BookingByBookIdOutput).Data)

	fmt.Println("TEST 5")
	if err != nil {
		baseOutput.Code = 500
	}

	output = outputs.BookingByBookIdOutput{
		BaseOutput: baseOutput,
		Data:       output.(outputs.BookingByBookIdOutput).Data,
	}
	return 201, output
}

func DeleteBookingByBookingId(ctx context.Context, bookID string) (int, interface{}) {
	db := configs.DB
	var booking database.Bookings
	err := db.Table("bookings").WithContext(ctx).First(&booking, "bookings.id = ?", bookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			output := outputs.NotFoundOutput{
				Code:    404,
				Message: fmt.Sprintf("Not Found: Booking %s not found", bookID),
			}
			return 404, output
		} else {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error",
			}
			return 500, output
		}
	}
	output := outputs.BaseOutput{
		Code:    201,
		Message: fmt.Sprintf("Success: Booking %s has been deleted", bookID),
	}
	return 201, output
}

func FindBookingByMentorID(mentorId string, ctx context.Context) (int, interface{}) {
	db := configs.GetDB()
	var listBooking []responses.BookingData = nil
	err := db.Table("bookings b").
		WithContext(ctx).
		Select("b.id AS booking_id, u.id AS user_id, u.username AS user_name, ts.mentorId AS mentor_id, ts.mentorName AS mentor_name, c.id AS course_id, c.name AS course_name, c.detail AS course_detail, i.id AS invoice_id, i.payment_method AS invoice_payment_method, i.payment_name AS invoice_payment_name, i.payment_status AS invoice_payment_status, i.payment_amount AS invoice_payment_amount, i.payment_fee AS invoice_payment_fee, i.payment_total AS invoice_payment_total, b.date AS booking_date, b.location AS location").
		Joins("JOIN users u ON b.user_id = u.id").
		Joins("JOIN courses c ON b.course_id = c.id").
		Joins("JOIN invoices i ON b.invoice_id = i.id").
		Joins(`JOIN (SELECT b2.id AS bookingID, m2.id AS mentorID, u2.username AS mentorName FROM bookings b2
			JOIN mentors m2 ON b2.mentor_id = m2.id
			JOIN users u2 ON u2.id = m2.user_id
			WHERE b2.mentor_id = ?) AS ts ON b.id = ts.bookingID`, mentorId).
		Scan(&listBooking).
		Error

	if err != nil {
		fmt.Println("Error:", err.Error())
		log.Println("Error_Log:", err.Error())
		return 500, gin.H{"message": "Internal Server Error"}
	}

	output := outputs.BookingByUserIdOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: List of bookings has been retrieved",
		},
		Data: DTO.ToBookingResponses(listBooking),
	}
	return 200, output
}
