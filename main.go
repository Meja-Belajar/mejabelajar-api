package main

import (
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/routers"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
}

// @BasePath /api/v1
func main() {
	// configs.DB.AutoMigrate(&database.Users{}, &database.Mentors{}, &database.Courses{}, &database.MentorCourses{}, &database.MentorReviews{}, &database.Invoices{}, &database.Bookings{})
	configs.DB.AutoMigrate(&database.Users{}, &database.Mentors{}, &database.Courses{}, &database.Invoices{}, &database.Bookings{})
	r := routers.ConfigureRouter()
	r.Run(":3000")
}
