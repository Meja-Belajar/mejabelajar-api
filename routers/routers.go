package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/services"
	"github.com/meja_belajar/middlewares"
)

func ConfigureRouter() *gin.Engine {
	//@adding middleware
	router := gin.New()
	//server dapat memberikan izin kepada klien dari domain yang berbeda untuk mengakses sumber daya
	router.Use(middlewares.CORSMiddleware())
	//mencatat detail permintaan HTTP
	router.Use(gin.Logger())
	//menangani panic yang terjadi selama penanganan permintaan
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"Code": 404, "Message": "Page Not Found"})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Code": 200, "Message": "Welcome to Meja Belajar API"})
	})
	//group yang perlu auth
	auth := router.Group("api/v1/auth")
	auth.Use(middlewares.RequiredAuth())
	services.UserServiceAuth(auth)
	services.MentorServiceAuth(auth)

	//group basic
	base := router.Group("api/v1")
	services.BookingService(base)
	services.UserServiceBasic(base)
	services.MentorReviewService(base)
	services.NotificationService(base)
	services.SearchService(base)

	return router
}
