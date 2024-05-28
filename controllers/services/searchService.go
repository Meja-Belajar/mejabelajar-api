package services

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/helpers"
)

func SearchService(router *gin.RouterGroup) {
	router.GET("/search/:query", Search)
}

func Search(c *gin.Context) {
	query := c.Param("query")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	code, output := helpers.Search(ctx, query)
	c.JSON(code, output)
}
