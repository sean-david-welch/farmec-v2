package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type MachineController struct {
	machineService services.MachineService
}

func NewMachineController(machineService services.MachineService) *MachineController {
	return &MachineController{machineService: machineService}
}

func (controller *MachineController) GetMachines(context *gin.Context) {
	id := context.Param("id")

	machines, err := controller.machineService.GetMachines(id)
	if err != nil {
		log.Printf("Error getting machines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting machines"})
		return
	}

	context.JSON(http.StatusOK, machines)
}

func (controller *MachineController) GetMachineById(context *gin.Context) {
	id := context.Param("id")

	machine, err := controller.machineService.GetMachineById(id)
	if err != nil {
		log.Printf("Error getting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting machine"})
		return
	}

	context.JSON(http.StatusOK, machine)
}

func (controller *MachineController) CreateMachine(context *gin.Context) {
	var machine types.Machine

	if err := context.ShouldBindJSON(&machine); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	result, err := controller.machineService.CreateMachine(&machine)
	if err != nil {
		log.Printf("Error creating machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating machine", "details": err.Error()})
		return
	}

	response := gin.H{
		"machine":      machine,
		"PresignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (controller *MachineController) UpdateMachine(context *gin.Context) {
	id := context.Param("id")

	var machine types.Machine

	if err := context.ShouldBindJSON(&machine); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	result, err := controller.machineService.UpdateMachine(id, &machine)
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

func (controller *MachineController) DeleteMachine(context *gin.Context) {
	id := context.Param("id")

	if err := controller.machineService.DeleteMachine(id); err != nil {
		log.Printf("Error deleting machine: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting machine", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Machine deleted successfully", "id": id})
}
