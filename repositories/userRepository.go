package repositories

import (
	"log"
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"gorm.io/gorm"
)

func InsertUser(user database.Users) error {
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)

	err = db.Create(&user).Error
	if ctx.Err() == context.DeadlineExceeded {
		return ctx.Err()
	}

	return err
}

func FindUserByEmail(email string) (database.Users, error) {
	var user database.Users
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return user, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := configs.GetDB().WithContext(ctx)
	err = db.
		Table("users").
		Where("email = ?", email).
		First(&user).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return user, ctx.Err()
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}

func FindUserByUserID(userID string) (database.Users, error) {
	log.Print("FindUserByUserID")
	var user database.Users
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return user, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)
	err = db.
		Table("users").
		Where("id = ?", userID).
		First(&user).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return user, ctx.Err()
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}
