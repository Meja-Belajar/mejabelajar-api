package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/meja_belajar/models/DTO"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/repositories"
)

func MakeNotification(ctx context.Context, booking responses.BookingResponseDTO) error {
	notification := database.Notification{}
	notification.Title = booking.Course.Name
	notification.UserId = booking.User.ID
	Description := "Booking Data<br/>Booking ID:" + booking.ID.String() +
		"<br/>Booking Date: " + booking.BookingDate.Local().String() +
		"<br/>Location: " + booking.BookingLocation +
		"<br/>User ID: " + booking.User.ID.String() +
		"<br/>User Name: " + booking.User.Name +
		"<br/>Mentor Name: " + booking.Mentor.Name +
		"<br/>Course Name: " + booking.Course.Name +
		"<br/>Course Detail: " + booking.Course.Detail
	notification.Description = Description

	return repositories.CreateNotification(ctx, notification)
}

func GetNotifications(ctx context.Context, userId string) (int, interface{}) {
	notifications, err := repositories.GetNotifications(ctx, userId)

	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Code: 500, Message: "Internal Server Error"}
	}

	if len(notifications) == 0 {
		return 404, outputs.NotFoundOutput{Code: 404, Message: "Notification Not Found"}
	}
	
	log.Println(notifications);

	output := outputs.WebResponse{}
	output.BaseOutput = outputs.BaseOutput{Code: 200, Message: "Success Get Notifications"}

	output.Data = DTO.ToNotificationResponses(notifications)
	return 200, output
}