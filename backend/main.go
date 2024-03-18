package main

import (
	// "github.com/meja_belajar/controllers"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&database.Users{}, &database.Mentors{}, &database.Courses{}, &database.MentorCourses{}, &database.Bookings{})

	// r := gin.Default()
	// r.POST("/posts", controllers.PostsCreate)
	// r.Run()
}