package services

import (
	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/requests"
)

func BookingService(router *gin.RouterGroup) {
	router.GET("/bookings/:userID", GetBookingByUserID)
	router.GET("/booking/:bookingID", GetBookingByBookingID)
	router.POST("/bookings", CreateBooking)
	router.DELETE("/booking/:bookingID", DeleteBooking)
}

func GetBookingByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	code, output := helpers.FindBookingByUserID(userID)
	ctx.JSON(code, output)
}

func GetBookingByBookingID(ctx *gin.Context) {
	BookingID := ctx.Param("bookingID")
	code, output := helpers.FindBookingByBookingID(BookingID)
	ctx.JSON(code, output)
}

func CreateBooking(ctx *gin.Context) {
	var requestData requests.NewBookingRequestDTO
	err := ctx.BindJSON(&requestData)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	code, output := helpers.CreateBooking(requestData)
	ctx.JSON(code, output)
}

func DeleteBooking(ctx *gin.Context) {
	BookingID := ctx.Param("bookingID")
	code, output := helpers.DeleteBookingByBookingId(BookingID)
	ctx.JSON(code, output)
}
