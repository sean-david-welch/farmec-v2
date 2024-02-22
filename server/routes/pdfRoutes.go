package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitPdfRenderer(router *gin.Engine, adminMiddleware *middleware.AdminMiddleware) {
	pdfService := services.NewPdfService()
	pdfController := controllers.NewPdfController(pdfService)

	PdfRenderRoutes(router, pdfController, adminMiddleware)
}

func PdfRenderRoutes(router *gin.Engine, pdfController *controllers.PdfController, adminMiddleware *middleware.AdminMiddleware) {
	pdfGroup := router.Group("/api/pdf")
	// protected := pdfGroup.Group("").Use(adminMiddleware.Middleware())
	{
		pdfGroup.POST("/registration", pdfController.RenderRegistrationPdf)
		pdfGroup.POST("/warranty", pdfController.RenderWarrantyClaimPdf)
	}
}
