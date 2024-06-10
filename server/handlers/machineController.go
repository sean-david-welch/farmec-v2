package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type MachineHandler struct {
	machineService services.MachineService
}

func NewMachineHandler(machineService services.MachineService) *MachineHandler {
	return &MachineHandler{machineService: machineService}
}

func (handler *MachineHandler) GetMachines(context *gin.Context) {
	id := context.Param("id")

	machines, err := handler.machineService.GetMachines(id)
	if err != nil {
		log.Printf("Error getting machines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting machines"})
		return
	}

	context.JSON(http.StatusOK, machines)
}

func (handler *MachineHandler) GetMachineById(context *gin.Context) {
	id := context.Param("id")

	machine, err := handler.machineService.GetMachineById(id)
	if err != nil {
		log.Printf("Error getting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting machine"})
		return
	}

	context.JSON(http.StatusOK, machine)
}

func (handler *MachineHandler) CreateMachine(context *gin.Context) {
	var machine types.Machine

	if err := context.ShouldBindJSON(&machine); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if machine.SupplierID == "" || machine.SupplierID == "null" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "SupplierID cannot be empty"})
		return
	}

	result, err := handler.machineService.CreateMachine(&machine)
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

	result, err := handler.machineService.UpdateMachine(id, &machine)
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
	id := context.Param("id")

	if err := handler.machineService.DeleteMachine(id); err != nil {
		log.Printf("Error deleting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Machine deleted successfully", "id": id})
}
