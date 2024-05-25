package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/requests"
)

func BookingService(router *gin.RouterGroup) {
	router.GET("/bookings/user/:userID", GetBookingByUserID)
	router.GET("/booking/:bookingID", GetBookingByBookingID)
	router.GET("/bookings/mentor/:mentorID", GetBookingByMentorID)

	router.POST("/booking", CreateBooking)
	router.DELETE("/booking/:bookingID", DeleteBooking)
}

func GetBookingByUserID(c *gin.Context) {
	userID := c.Param("userID")
	fmt.Println("First Service")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()
	code, output := helpers.FindBookingByUserID(userID, ctx)
	c.JSON(code, output)
}

func GetBookingByBookingID(c *gin.Context) {
	BookingID := c.Param("bookingID")
	code, output := helpers.FindBookingByBookingID(BookingID)
	c.JSON(code, output)
}

func GetBookingByMentorID(c *gin.Context) {
	MentorId := c.Param("mentorID")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()
	code, output := helpers.FindBookingByMentorID(MentorId, ctx)
	c.JSON(code, output)
}

func CreateBooking(c *gin.Context) {
	var requestData requests.NewBookingRequestDTO
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	
	log.Println(requestData);

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	code, output := helpers.CreateBooking(ctx, requestData)
	c.JSON(code, output)
}

func DeleteBooking(c *gin.Context) {
	BookingID := c.Param("bookingID")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	code, output := helpers.DeleteBookingByBookingId(ctx, BookingID)
	c.JSON(code, output)
}
