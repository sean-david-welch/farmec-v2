package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type PrivacyHandler struct {
	service services.PrivacyService
}

func NewPrivacyHandler(service services.PrivacyService) *PrivacyHandler {
	return &PrivacyHandler{service: service}
}

func (handler *PrivacyHandler) GetPrivacys(context *gin.Context) {
	ctx := context.Request.Context()
	privacys, err := handler.service.GetPrivacys(ctx)
	if err != nil {
		log.Printf("error getting privacys: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting privacy"})
		return
	}

	context.JSON(http.StatusOK, privacys)
}

func (handler *PrivacyHandler) CreatePrivacy(context *gin.Context) {
	ctx := context.Request.Context()
	var privacy types.Privacy

	if err := context.ShouldBindJSON(&privacy); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbPrivacy := lib.DeserializePrivacy(privacy)
	if err := handler.service.CreatePrivacy(ctx, &dbPrivacy); err != nil {
		log.Printf("error while creating privacy: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating privacy", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, privacy)
}

func (handler *PrivacyHandler) UpdatePrivacy(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	var privacy types.Privacy
	if err := context.ShouldBindJSON(&privacy); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbPrivacy := lib.DeserializePrivacy(privacy)
	if err := handler.service.UpdatePrivacy(ctx, id, &dbPrivacy); err != nil {
		log.Printf("error while updating privacy: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating privacy", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, privacy)
}

func (handler *PrivacyHandler) DeletePrivacy(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeletePrivacy(ctx, id); err != nil {
		log.Printf("Error deleting privacy: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting privacy", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "privacy deleted successfully", "id": id})
}
