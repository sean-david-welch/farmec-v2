package routes

import (
	"database/sql"
	"embed"
	"github.com/sean-david-welch/farmec-v2/server/resources"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
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

	// default routes
	router.GET("/", func(c *gin.Context) {
		component := pages.Home([]types.Carousel{}, supplierCache.GetSuppliersFromContext(c))
		if err := component.Render(c.Request.Context(), c.Writer); err != nil {
			log.Printf("Error rendering view: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering view"})
			return
		}
		c.Header("Content-Type", "text/html; charset=utf-8")
	})

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Farmec Ireland's API Service.",
		})
	})

	// main routes
	SupplierRoutes(router, resourcesMap.SupplierHandler, authMiddleware)
	ViewRoutes(router, resourcesMap.ViewHandler, authMiddleware, supplierCache)
	CarouselRoutes(router, resourcesMap.CarouselHandler, authMiddleware)

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
