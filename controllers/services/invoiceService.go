package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
)

func InvoiceService(router *gin.RouterGroup) {
	router.GET("/invoices/user/:userID", GetInvoiceByUserID)
	router.GET("/invoice/:invoiceID", GetInvoiceByID)

	router.POST("/invoice/update", UpdateInvoiceStatus)
}


func GetInvoiceByUserID(c *gin.Context) {
	userID := c.Param("userID")
	fmt.Println("First Service")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()
	code, output := helpers.FindInvoiceByUserID(userID, ctx)
	c.JSON(code, output)
}

func GetInvoiceByID(c *gin.Context) {
	InvoiceID := c.Param("invoiceID")
	code, output := helpers.FindInvoiceByInvoiceID(InvoiceID)
	c.JSON(code, output)
}

func UpdateInvoiceStatus(c *gin.Context) {
	var UpdateInvoiceStatusRequestDTO requests.UpdateInvoiceStatusRequestDTO

	if err := c.ShouldBindJSON(&UpdateInvoiceStatusRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.UpdateInvoiceStatus(UpdateInvoiceStatusRequestDTO)
	c.JSON(code, output)
}
