package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (handler *EmployeeHandler) GetEmployees(context *gin.Context) {
	employees, err := handler.service.GetEmployees()
	if err != nil {
		log.Printf("error getting employees: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting employees"})
		return
	}

	context.JSON(http.StatusOK, employees)
}

func (handler *EmployeeHandler) CreateEmployee(context *gin.Context) {
	var employee types.Employee

	if err := context.ShouldBindJSON(&employee); err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := handler.service.CreateEmployee(&employee)
	if err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating employee", "details": err.Error()})
		return
	}

	response := gin.H{
		"employee":     employee,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusCreated, response)
}

func (handler *EmployeeHandler) UpdateEmployee(context *gin.Context) {
	id := context.Param("id")
	var employee types.Employee

	if err := context.ShouldBindJSON(&employee); err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	result, err := handler.service.UpdateEmployee(id, &employee)
	if err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating employee", "details": err.Error()})
		return
	}

	response := gin.H{
		"employee":     employee,
		"presignedUrl": result.PresignedUrl,
		"imageUrl":     result.ImageUrl,
	}

	context.JSON(http.StatusAccepted, response)
}

func (handler *EmployeeHandler) DeleteEmployee(context *gin.Context) {
	id := context.Param("id")

	if err := handler.service.DeleteEmployee(id); err != nil {
		log.Printf("Error deleting employee: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting employee", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully", "id": id})
}
