package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"gorm.io/gorm"
)

func FindCourseByID(courseID string) (database.Courses, error) {
	var courses database.Courses

	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return courses, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)

	err = db.
		Table("courses").
		Where("courses.id", courseID).
		Find(&courses).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return courses, ctx.Err()
	}

	//validasi not found
	if err == gorm.ErrRecordNotFound {
		return courses, err
	}

	if err != nil {
		return courses, err
	}
	return courses, nil
}
