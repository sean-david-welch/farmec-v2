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
