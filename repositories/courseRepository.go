package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"gorm.io/gorm"
)

// Get All Courses
func GetAllCourse(ID string) (database.Courses, error) {
	var course database.Courses
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return course, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)
	err = db.
		Table("courses").
		Find(&course).
		Error

	if ctx.Err() == context.DeadlineExceeded {
		return course, ctx.Err()
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return course, err
	}
	return course, nil
}