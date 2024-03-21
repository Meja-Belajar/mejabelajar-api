package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/controllers/services"

)
func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&database.Users{}, &database.Mentors{}, &database.Courses{}, &database.MentorCourses{}, &database.Bookings{})

	r := gin.Default()
	// r.POST("/posts", controllers.PostsCreate)
	services.UserService(r)
	r.Run(":3000")
}
