package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func PdfRenderRoutes(router *gin.Engine, handler *handlers.PdfHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	pdfGroup := router.Group("/api/pdf")
	protected := pdfGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("/registration", handler.RenderRegistrationPdf)
		protected.POST("/warranty", handler.RenderWarrantyClaimPdf)
	}
}
