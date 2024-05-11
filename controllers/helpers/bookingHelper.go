package helpers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/responses"
	"gorm.io/gorm"
)

func FindBookingByUserID(userID string) (int, interface{}) {
	db := configs.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

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
			Message: "Success: booking has been retrieved",
		},
		Data: booking,
	}
	return 200, output
}

func DeleteBookingByBookingId(bookID string) (int, interface{}) {
	db := configs.DB

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

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
		Code:    200,
		Message: fmt.Sprintf("Success: Booking %s has been deleted", bookID),
	}
	return 200, output
}
