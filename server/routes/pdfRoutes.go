package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitPdfRenderer(router *gin.Engine, adminMiddleware *middleware.AdminMiddleware) {
	pdfService := services.NewPdfService()
	pdfHandler := handlers.NewPdfHandler(pdfService)

	PdfRenderRoutes(router, pdfHandler, adminMiddleware)
}

func PdfRenderRoutes(router *gin.Engine, pdfHandler *handlers.PdfHandler, adminMiddleware *middleware.AdminMiddleware) {
	pdfGroup := router.Group("/api/pdf")
	// protected := pdfGroup.Group("").Use(adminMiddleware.Middleware())
	{
		pdfGroup.POST("/registration", pdfHandler.RenderRegistrationPdf)
		pdfGroup.POST("/warranty", pdfHandler.RenderWarrantyClaimPdf)
	}
}
