package DTO

import (
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/responses"
)

func ToNotificationResponse(notification database.Notification) responses.NotificationReponse {
	return responses.NotificationReponse{
		ID:          notification.ID,
		Title:       notification.Title,
		Description: notification.Description,
		CreatedAt:   notification.CreatedAt,
	}
}

func ToNotificationResponses(notifications []database.Notification) []responses.NotificationReponse {
	var notificationResponses []responses.NotificationReponse = nil
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, ToNotificationResponse(notification))
	}
	return notificationResponses
}
