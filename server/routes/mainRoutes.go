package routes

import (
	"database/sql"
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
	s3Client lib.S3Client, firebase *lib.Firebase, smtp *lib.SMTPClientImpl,
	adminMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache,
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

	// Supplier Module Resources
	SupplierRoutes(router, supplierHandler, adminMiddleware)
	InitParts(router, database, s3Client, adminMiddleware)
	InitVideos(router, database, secrets, adminMiddleware)
	InitProduct(router, database, s3Client, adminMiddleware)
	InitMachines(router, database, s3Client, adminMiddleware)

	// About Module Resources
	InitTimelines(router, database, adminMiddleware)
	InitializeEmployee(router, database, s3Client, adminMiddleware)

	// Blog Module Resources
	InitExhibitions(router, database, adminMiddleware, supplierCache)
	InitBlogs(router, database, s3Client, adminMiddleware)

	// Misc Resources
	InitWarranty(router, database, adminMiddleware, smtp)
	InitRegistrations(router, database, adminMiddleware, smtp)
	InitLineItems(router, database, s3Client, adminMiddleware)
	InitCarousel(router, database, s3Client, adminMiddleware)

	// Util Resources
	InitContact(router, smtp)
	InitCheckout(router, database, secrets)
	InitPdfRenderer(router, adminMiddleware)
	InitAuth(router, firebase, adminMiddleware)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})
}
