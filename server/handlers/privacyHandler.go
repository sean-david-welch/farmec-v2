package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyHandler struct {
	service services.PrivacyService
}

func NewPrivacyHandler(service services.PrivacyService) *PrivacyHandler {
	return &PrivacyHandler{service: service}
}

func (handler *PrivacyHandler) GetPrivacys(context *gin.Context) {
	privacys, err := handler.service.GetPrivacys()
	if err != nil {
		log.Printf("error getting privacys: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting privacy"})
		return
	}

	context.JSON(http.StatusOK, privacys)
}

func (handler *PrivacyHandler) CreatePrivacy(context *gin.Context) {
	var privacy types.Privacy

	if err := context.ShouldBindJSON(&privacy); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := handler.service.CreatePrivacy(&privacy); err != nil {
		log.Printf("error while creating privacy: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating privacy", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, privacy)
}

func (handler *PrivacyHandler) UpdatePrivacy(context *gin.Context) {
	id := context.Param("id")
	var privacy types.Privacy

	if err := context.ShouldBindJSON(&privacy); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := handler.service.UpdatePrivacy(id, &privacy); err != nil {
		log.Printf("error while updating privacy: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating privacy", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, privacy)
}

func (handler *PrivacyHandler) DeletePrivacy(context *gin.Context) {
	id := context.Param("id")

	if err := handler.service.DeletePrivacy(id); err != nil {
		log.Printf("Error deleting privacy: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting privacy", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "privacy deleted successfully", "id": id})
}
