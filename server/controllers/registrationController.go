package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type RegistrationController struct {
	service *services.RegistrationService
}

func NewRegistrationController(service *services.RegistrationService) *RegistrationController {
	return &RegistrationController{service: service}
}

func(controller *RegistrationController) GetRegistrations(context *gin.Context) {
	registrations, err := controller.service.GetRegistrations(); if err != nil {
		log.Printf("error occurred while getting registrations: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting registration"})
		return
	}

	context.JSON(http.StatusOK, registrations)
} 

func(controller *RegistrationController) GetRegistrationById(context *gin.Context) {
	id := context.Param("id")

	registration, err := controller.service.GetRegistrationById(id); if err != nil {
		log.Printf("error occurred while getting registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting registration"})
		return
	}

	context.JSON(http.StatusOK, registration)
}

func(controller *RegistrationController) CreateRegistration(context *gin.Context) {
	var registration *models.MachineRegistration

	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error when creating registration: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while creating registration -  bad request"})
		return
	}

	if err := controller.service.CreateRegistration(registration); err != nil {
		log.Printf("error when creating registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when creating registration"})
	}

	context.JSON(http.StatusCreated, registration)
}

func(controller *RegistrationController) UpdateRegistration(context *gin.Context) {
	id := context.Param("id")
	var registration *models.MachineRegistration

	if err := context.ShouldBindJSON(&registration); err != nil {
		log.Printf("error when updating registration: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating registration -  bad request"})
		return
	}

	if err := controller.service.UpdateRegistration(id, registration); err != nil {
		log.Printf("error when updating registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error when updating registration"})
	}

	context.JSON(http.StatusAccepted, registration)
}

func(controller *RegistrationController) DeleteRegistration(context *gin.Context) {
	id := context.Param("id")

	if err := controller.service.DeleteRegistration(id); err != nil {
		log.Printf("error occurred while deleting registration: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting registration"})
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "registration deleted successfully", "id": id})
}

