package services

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/helpers"
)

func NotificationService(router *gin.RouterGroup) {
	router.GET("/bookings/notification/:userID", GetNotifications)
}

func GetNotifications(c *gin.Context) {
	userId := c.Param("userID")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	code, output := helpers.GetNotifications(ctx, userId)
	c.JSON(code, output)
}
