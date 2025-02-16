package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
	"net/http"
)

type AboutHandler struct {
	employeeService services.EmployeeService
	timelineService services.TimelineService
	authMiddleware  *middleware.AuthMiddlewareImpl
	supplierCache   *middleware.SupplierCache
}

func NewAboutHandler(employeeService services.EmployeeService, timelineService services.TimelineService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *AboutHandler {
	return &AboutHandler{employeeService: employeeService, timelineService: timelineService, authMiddleware: authMiddleware, supplierCache: supplierCache}
}

func (handler *AboutHandler) AboutView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	employees, err := handler.employeeService.GetEmployees(request)
	if err != nil {
		log.Printf("Error getting employees: %v", err)
	}
	timelines, err := handler.timelineService.GetTimelines(request)
	if err != nil {
		log.Printf("Error getting timelines: %v", err)
	}

	isError := err != nil
	component := pages.About(isAdmin, isError, employees, timelines, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering view: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering view"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}

// Employee Hanlders
func (handler *AboutHandler) GetEmployees(context *gin.Context) {
	request := context.Request.Context()
	employees, err := handler.employeeService.GetEmployees(request)
	if err != nil {
		log.Printf("error getting employees: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting employees"})
		return
	}

	context.JSON(http.StatusOK, employees)
}

func (handler *AboutHandler) CreateEmployee(context *gin.Context) {
	request := context.Request.Context()
	var employee types.Employee

	if err := context.ShouldBindJSON(&employee); err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbEmployee := lib.DeserializeEmployee(employee)
	result, err := handler.employeeService.CreateEmployee(request, &dbEmployee)
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

func (handler *AboutHandler) UpdateEmployee(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	var employee types.Employee
	if err := context.ShouldBindJSON(&employee); err != nil {
		log.Printf("Error creating employee: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbEmployee := lib.DeserializeEmployee(employee)
	result, err := handler.employeeService.UpdateEmployee(request, id, &dbEmployee)
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

func (handler *AboutHandler) DeleteEmployee(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.employeeService.DeleteEmployee(request, id); err != nil {
		log.Printf("Error deleting employee: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting employee", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully", "id": id})
}

// Timeline Hanlders
func (handler *AboutHandler) GetTimelines(context *gin.Context) {
	request := context.Request.Context()
	timelines, err := handler.timelineService.GetTimelines(request)
	if err != nil {
		log.Printf("error getting timelines: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting timeline"})
		return
	}

	context.JSON(http.StatusOK, timelines)
}

func (handler *AboutHandler) CreateTimeline(context *gin.Context) {
	request := context.Request.Context()
	var timeline types.Timeline

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbTimeline := lib.DeserializeTimeline(timeline)
	if err := handler.timelineService.CreateTimeline(request, &dbTimeline); err != nil {
		log.Printf("error while creating timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, timeline)
}

func (handler *AboutHandler) UpdateTimeline(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")
	var timeline types.Timeline

	if err := context.ShouldBindJSON(&timeline); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbTimeline := lib.DeserializeTimeline(timeline)
	if err := handler.timelineService.UpdateTimeline(request, id, &dbTimeline); err != nil {
		log.Printf("error while updating timeline: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, timeline)
}

func (handler *AboutHandler) DeleteTimeline(context *gin.Context) {
	request := context.Request.Context()
	id := context.Param("id")

	if err := handler.timelineService.DeleteTimeline(request, id); err != nil {
		log.Printf("Error deleting timeline: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting timeline", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "timeline deleted successfully", "id": id})
}
