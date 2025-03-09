package routes

import (
	"database/sql"
	"embed"
	"github.com/sean-david-welch/farmec-v2/server/resources"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func InitRoutes(
	router *gin.Engine, database *sql.DB, secrets *lib.Secrets,
	s3Client lib.S3Client, files embed.FS, firebase *lib.Firebase, emailClient *lib.EmailClientImpl,
	authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache,
) {
	resourcesMap := resources.NewResources(database, secrets, firebase, files, s3Client, emailClient, authMiddleware, supplierCache)
	// define supplier middleware
	router.Use(middleware.WithSupplierCache(supplierCache, resourcesMap.SupplierService))

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Farmec Ireland's API Service.",
		})
	})

	// main routes
	SupplierRoutes(router, resourcesMap.SupplierHandler, authMiddleware)
	ViewRoutes(router, resourcesMap.ViewHandler, authMiddleware)
	CarouselRoutes(router, resourcesMap.CarouselHandler, authMiddleware)

	// Supplier Module Resources
	PartsRoutes(router, resourcesMap.PartsHandler, authMiddleware)
	VideoRoutes(router, resourcesMap.VideoHandler, authMiddleware)
	ProductRoutes(router, resourcesMap.ProductHandler, authMiddleware)
	MachineRoutes(router, resourcesMap.MachineHandler, authMiddleware)

	// About Module Resources
	AboutRoutes(router, resourcesMap.AboutHandler, authMiddleware)

	// Blog Module Resources
	BlogRoutes(router, resourcesMap.BlogHandler, authMiddleware)

	// Misc Resources
	WarrantyRoutes(router, resourcesMap.WarrantyHandler, authMiddleware)
	RegistrationRoutes(router, resourcesMap.RegistrationHandler, authMiddleware)
	LineItemRoutes(router, resourcesMap.LineItemHandler, authMiddleware)

	// Util Resources
	CheckoutRoutes(router, resourcesMap.CheckoutHandler)
	PdfRenderRoutes(router, resourcesMap.PdfHandler, authMiddleware)
	AuthRoutes(router, resourcesMap.AuthHandler, authMiddleware)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})
}
