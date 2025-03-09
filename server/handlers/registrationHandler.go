package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"github.com/sean-david-welch/farmec-v2/server/views/pages/details"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type RegistrationHandler struct {
	service        services.RegistrationService
	authMiddleware *middleware.AuthMiddlewareImpl
	supplierCache  *middleware.SupplierCache
}

func NewRegistrationHandler(service services.RegistrationService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *RegistrationHandler {
	return &RegistrationHandler{service: service, authMiddleware: authMiddleware, supplierCache: supplierCache}
}

func (handler *RegistrationHandler) RegistrationsView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	isAuthenticated := handler.authMiddleware.GetIsAuthenticated(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	registrations, err := handler.service.GetRegistrations(request)
	if err != nil {
		log.Printf("Error getting registrations: %v\n", err)
	}
	component := pages.Registrations(isAdmin, isAuthenticated, registrations, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering registrations: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering the page"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

func (handler *RegistrationHandler) RegistrationsDetailView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	id := context.Param("id")
	registration, err := handler.service.GetRegistrationById(request, id)
	if err != nil {
		log.Printf("Error getting registration: %v\n", err)
	}

	isError := err != nil
	component := details.ReigstrationDetail(isAdmin, isError, *registration, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering registrations: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering the page"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

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
