package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ContactController struct {
	service services.ContactService
}

func NewContactController(service services.ContactService) *ContactController {
	return &ContactController{service: service}
}

func (controller *ContactController) SendEmail(context *gin.Context) {
	var data *types.EmailData

	if err := context.ShouldBindJSON(&data); err != nil {
		log.Printf("error in data - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error in request format"})
		return
	}

	if err := controller.service.SendEmail(data); err != nil {
		log.Printf("internal server error: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}
