package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitPdfRenderer(router *gin.Engine, adminMiddleware *middleware.AuthMiddlewareImpl) {
	service := services.NewPdfService()
	handler := handlers.NewPdfHandler(service)

	PdfRenderRoutes(router, handler, adminMiddleware)
}

func PdfRenderRoutes(router *gin.Engine, handler *handlers.PdfHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	pdfGroup := router.Group("/api/pdf")
	protected := pdfGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.POST("/registration", handler.RenderRegistrationPdf)
		protected.POST("/warranty", handler.RenderWarrantyClaimPdf)
	}
}
