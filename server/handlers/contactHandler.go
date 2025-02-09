package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactHandler struct {
	service services.ContactService
}

func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

func (handler *ContactHandler) SendEmail(context *gin.Context) {
	var data *types.EmailData

	if err := context.ShouldBindJSON(&data); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	if err := handler.service.SendContactEmail(data); err != nil {
		log.Printf("Failed to send email: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to send email",
			"message": "Please try again later or contact support",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully",
	})
}
