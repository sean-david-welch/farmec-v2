package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
	"net/http"
)

type AboutHandler struct {
	employeeService services.EmployeeServiceImpl
	timelineService services.TimelineServiceImpl
	adminMiddleware *middleware.AuthMiddlewareImpl
	supplierCache   *middleware.SupplierCache
}

func NewAboutHandler(employeeService services.EmployeeServiceImpl, timelineService services.TimelineServiceImpl, adminMiddleware *middleware.AuthMiddlewareImpl) *AboutHandler {
	return &AboutHandler{employeeService: employeeService, timelineService: timelineService, adminMiddleware: adminMiddleware}
}

func (handler *AboutHandler) AboutView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.adminMiddleware.GetIsAdmin(context)
	suppliers := middleware.GetSuppliersFromContext(context)

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
