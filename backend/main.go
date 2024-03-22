package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/controllers/services"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/utils"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&database.Users{}, &database.Mentors{}, &database.Courses{}, &database.MentorCourses{}, &database.Bookings{}, &database.MentorReviews{}, &database.Invoices{})
	hashedPassword, err := utils.HashPassword("P@ssw0rd!")
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Println("Test hashed password ", hashedPassword)
	r := gin.Default()
	// r.POST("/posts", controllers.PostsCreate)
	services.UserService(r)
	r.Run(":3000")
}
