package routes

import (
	"database/sql"
	"embed"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
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
	// instantiate supplier resouces and middleware
	supplierRepository := repository.NewSupplierRepo(database)
	supplierService := services.NewSupplierService(supplierRepository, s3Client, "Suppliers")
	supplierHandler := handlers.NewSupplierContoller(supplierService)

	// define supplier middleware
	router.Use(middleware.WithSupplierCache(supplierCache, supplierService))

	// default routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Farmec Ireland's API Service.",
		})
	})

	router.GET("/api", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Redirect /api requests without a valid route to /api root
	router.NoRoute(func(c *gin.Context) {
		// Return 404 for non-existent API routes
		if c.Request.URL.Path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "API route not found",
			})
			return
		}
		// Handle other non-API requests (which should go to the frontend)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	// main routes
	SupplierRoutes(router, supplierHandler, authMiddleware)

	// Supplier Module Resources
	InitParts(router, database, s3Client, authMiddleware)
	InitVideos(router, database, secrets, authMiddleware)
	InitProduct(router, database, s3Client, authMiddleware)
	InitMachines(router, database, s3Client, authMiddleware)

	// About Module Resources
	InitAbout(router, database, s3Client, authMiddleware, supplierCache)

	// Blog Module Resources
	InitBlogs(router, database, s3Client, authMiddleware, supplierCache)

	// Misc Resources
	InitWarranty(router, database, authMiddleware, emailClient)
	InitRegistrations(router, database, authMiddleware, emailClient)
	InitLineItems(router, database, s3Client, authMiddleware)
	InitCarousel(router, database, s3Client, authMiddleware, supplierCache)

	// Util Resources
	InitContact(router, emailClient)
	InitCheckout(router, database, secrets)
	InitPdfRenderer(router, authMiddleware, files)
	InitAuth(router, firebase, authMiddleware)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})
}
