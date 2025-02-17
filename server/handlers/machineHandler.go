package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type MachineHandler struct {
	service        services.MachineService
	authMiddleware *middleware.AuthMiddlewareImpl
	supplierCache  *middleware.SupplierCache
}

func NewMachineHandler(service services.MachineService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *MachineHandler {
	return &MachineHandler{service: service, authMiddleware: authMiddleware, supplierCache: supplierCache}
}

//	func (handler *MachineHandler) MachineView(context *gin.Context) {
//		request := context.Request.Context()
//		isAdmin := handler.authMiddleware.GetIsAdmin(context)
//		id := context.Param("id")
//
//		machines, err := handler.service.GetMachines(context, id)
//		if err != nil {
//			log.Printf("Error getting machines: %v\n", err)
//		}
//
//		isError := err != nil
//		component := pag
//	}
func (handler *MachineHandler) GetMachines(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	machines, err := handler.service.GetMachines(request, id)
	if err != nil {
		log.Printf("Error getting machines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting machines"})
		return
	}

	context.JSON(http.StatusOK, machines)
}

func (handler *MachineHandler) GetMachineById(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	machine, err := handler.service.GetMachineById(request, id)
	if err != nil {
		log.Printf("Error getting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting machine"})
		return
	}

	context.JSON(http.StatusOK, machine)
}

func (handler *MachineHandler) CreateMachine(context *gin.Context) {
	request := context.Request.Context()
	var machine types.Machine

	if err := context.ShouldBindJSON(&machine); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	if machine.SupplierID == "" || machine.SupplierID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	dbMachine := lib.DeserializeMachine(machine)
	result, err := handler.service.CreateMachine(request, &dbMachine)
	if err != nil {
		log.Printf("Error creating machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating machine", "details": err.Error()})
		return
	}

	response := gin.H{
		"machine":      machine,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *MachineHandler) UpdateMachine(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var machine types.Machine
	if err := context.ShouldBindJSON(&machine); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	if machine.SupplierID == "" || machine.SupplierID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	dbMachine := lib.DeserializeMachine(machine)
	result, err := handler.service.UpdateMachine(request, id, &dbMachine)
	if err != nil {
		log.Printf("Error updating machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating machine", "details": err.Error()})
		return
	}

	response := gin.H{
		"machine":      machine,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *MachineHandler) DeleteMachine(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteMachine(request, id); err != nil {
		log.Printf("Error deleting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Machine deleted successfully", "id": id})
}
