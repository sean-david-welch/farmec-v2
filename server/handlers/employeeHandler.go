package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (handler *EmployeeHandler) GetEmployees(context *gin.Context) {
	request := context.Request.Context()
	employees, err := handler.service.GetEmployees(request)
	if err != nil {
		log.Printf("error getting employees: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting employees"})
		return
	}

	context.JSON(http.StatusOK, employees)
}

func (handler *EmployeeHandler) CreateEmployee(context *gin.Context) {
	request := context.Request.Context()
	var employee types.Employee

	if err := context.ShouldBindJSON(&employee); err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbEmployee := lib.DeserializeEmployee(employee)
	result, err := handler.service.CreateEmployee(request, &dbEmployee)
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
	request := context.Request.Context()
	id := context.Param("id")

	var employee types.Employee
	if err := context.ShouldBindJSON(&employee); err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbEmployee := lib.DeserializeEmployee(employee)
	result, err := handler.service.UpdateEmployee(request, id, &dbEmployee)
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
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteEmployee(request, id); err != nil {
		log.Printf("Error deleting employee: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting employee", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully", "id": id})
}
