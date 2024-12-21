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
		log.Printf("error in data - bad request: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error in request format"})
		return
	}
	go func() {
		err := handler.service.SendContactEmail(data)
		if err != nil {
			return
		}
	}()

	context.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}
