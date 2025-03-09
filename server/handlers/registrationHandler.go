package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type RegistrationHandler struct {
	service services.RegistrationService
}

func NewRegistrationHandler(service services.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{service: service}
}

//func (handler *RegistrationHandler) (context *gin.Context) {}

func (handler *RegistrationHandler) GetRegistrations(context *gin.Context) {
	request := context.Request.Context()
	registrations, err := handler.service.GetRegistrations(request)
	if err != nil {
		log.Printf("error occurred while getting registrations: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting registration"})
		return
	}

	context.JSON(http.StatusOK, registrations)
}

func (handler *RegistrationHandler) GetRegistrationById(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	registration, err := handler.service.GetRegistrationById(request, id)
	if err != nil {
		log.Printf("error occurred while getting registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting registration"})
		return
	}

	context.JSON(http.StatusOK, registration)
}

func (handler *RegistrationHandler) CreateRegistration(context *gin.Context) {
	request := context.Request.Context()
	var registration types.MachineRegistration
	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error when creating registration: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating registration -  bad request"})
		return
	}

	dbRegistration := lib.DeserializeMachineRegistration(registration)
	if err := handler.service.CreateRegistration(request, &dbRegistration); err != nil {
		log.Printf("error when creating registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when creating registration"})
	}

	context.JSON(http.StatusCreated, registration)
}

func (handler *RegistrationHandler) UpdateRegistration(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var registration types.MachineRegistration
	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error when updating registration: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating registration -  bad request"})
		return
	}

	dbRegistration := lib.DeserializeMachineRegistration(registration)
	if err := handler.service.UpdateRegistration(request, id, &dbRegistration); err != nil {
		log.Printf("error when updating registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when updating registration"})
	}

	context.JSON(http.StatusAccepted, registration)
}

func (handler *RegistrationHandler) DeleteRegistration(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteRegistration(request, id); err != nil {
		log.Printf("error occurred while deleting registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting registration"})
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "registration deleted successfully", "id": id})
}
