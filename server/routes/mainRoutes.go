package routes

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func InitRoutes(
	reactApp embed.FS, router *gin.Engine, database *sql.DB, secrets *lib.Secrets, s3Client lib.S3Client,
	adminMiddleware *middleware.AdminMiddleware, authMiddleware *middleware.AuthMiddleware, firebase *lib.Firebase, emailClient *lib.EmailClientImpl,
) {
	reactFS, err := fs.Sub(reactApp, "../client/dist")
	if err != nil {
		log.Fatal("Failed to create sub filesystem:", err)
	}
	// Serve static files from the embedded filesystem
	router.StaticFS("/assets", http.FS(reactFS))
	// Try to serve favicon and manifest from embedded files
	router.GET("/favicon.svg", func(c *gin.Context) {
		c.FileFromFS("favicon.svg", http.FS(reactFS))
	})
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Farmec Ireland's API Service.",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		if c.Request.URL.Path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "API route not found",
			})
			return
		}
		c.FileFromFS("index.html", http.FS(reactFS))
	})
	// Supplier Module Resources
	InitParts(router, database, s3Client, adminMiddleware)
	InitVideos(router, database, secrets, adminMiddleware)
	InitProduct(router, database, s3Client, adminMiddleware)
	InitMachines(router, database, s3Client, adminMiddleware)
	InitSuppliers(router, database, s3Client, adminMiddleware)
	// About Module Resources
	InitTerms(router, database, adminMiddleware)
	InitPrivacy(router, database, adminMiddleware)
	InitTimelines(router, database, adminMiddleware)
	InitializeEmployee(router, database, s3Client, adminMiddleware)
	// Blog Module Resources
	InitExhibitions(router, database, adminMiddleware)
	InitBlogs(router, database, s3Client, adminMiddleware)
	// Misc Resources
	InitWarranty(router, database, authMiddleware, emailClient)
	InitRegistrations(router, database, authMiddleware, emailClient)
	InitLineItems(router, database, s3Client, adminMiddleware)
	InitCarousel(router, database, s3Client, adminMiddleware)
	// Util Resources
	InitContact(router, emailClient)
	InitCheckout(router, database, secrets)
	InitPdfRenderer(router, adminMiddleware)
	InitAuth(router, firebase, adminMiddleware)
}
