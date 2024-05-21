package repositories

import (
	"context"
	"fmt"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
)

func CreateNotification(ctx context.Context, notification database.Notification) error {
	db := configs.GetDB()
	fmt.Println("Testing")
	err := db.WithContext(ctx).
		Model(&database.Notification{}).
		Create(&notification).
		Error

	fmt.Println(err)
	return err
}

func GetNotifications(ctx context.Context, userId string) ([]database.Notification, error) {
	var notifications []database.Notification = nil
	db := configs.GetDB()
	err := db.WithContext(ctx).
		Model(&database.Notification{}).
		Where("user_id = ?", userId).
		Scan(&notifications).
		Error
	return notifications, err
}
