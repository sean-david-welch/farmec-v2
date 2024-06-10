package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type RegistrationHandler struct {
	service services.RegistrationService
}

func NewRegistrationHandler(service services.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{service: service}
}

func (handler *RegistrationHandler) GetRegistrations(context *gin.Context) {
	registrations, err := handler.service.GetRegistrations()
	if err != nil {
		log.Printf("error occurred while getting registrations: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting registration"})
		return
	}

	context.JSON(http.StatusOK, registrations)
}

func (handler *RegistrationHandler) GetRegistrationById(context *gin.Context) {
	id := context.Param("id")

	registration, err := handler.service.GetRegistrationById(id)
	if err != nil {
		log.Printf("error occurred while getting registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting registration"})
		return
	}

	context.JSON(http.StatusOK, registration)
}

func (handler *RegistrationHandler) CreateRegistration(context *gin.Context) {
	var registration *types.MachineRegistration

	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error when creating registration: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating registration -  bad request"})
		return
	}

	if err := handler.service.CreateRegistration(registration); err != nil {
		log.Printf("error when creating registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when creating registration"})
	}

	context.JSON(http.StatusCreated, registration)
}

func (handler *RegistrationHandler) UpdateRegistration(context *gin.Context) {
	id := context.Param("id")
	var registration *types.MachineRegistration

	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error when updating registration: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating registration -  bad request"})
		return
	}

	if err := handler.service.UpdateRegistration(id, registration); err != nil {
		log.Printf("error when updating registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when updating registration"})
	}

	context.JSON(http.StatusAccepted, registration)
}

func (handler *RegistrationHandler) DeleteRegistration(context *gin.Context) {
	id := context.Param("id")

	if err := handler.service.DeleteRegistration(id); err != nil {
		log.Printf("error occurred while deleting registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting registration"})
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "registration deleted successfully", "id": id})
}
