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
	configs.DB.AutoMigrate(&database.Users{}, &database.Mentors{}, &database.Courses{}, &database.MentorCourses{}, &database.Bookings{}, &database.MentorReviews{}, &database.Invoices{})
	//r := gin.Default()
	//services.UserService(r)
	r := routers.ConfigureRouter()
	r.Run(":3000")
}
