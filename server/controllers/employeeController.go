package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func(controller *EmployeeController) GetEmployees(context *gin.Context) {
	employees, err := controller.service.GetEmployees(); if err != nil {
		log.Printf("error getting employees: %w", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting employees"})
		return
	}
}