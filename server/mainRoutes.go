package main

import (
	"database/sql"
	"embed"
	"github.com/sean-david-welch/farmec-v2/server/routes"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func InitRoutes(
	reactApp embed.FS, router *gin.Engine, database *sql.DB, secrets *lib.Secrets, s3Client lib.S3Client,
	adminMiddleware *middleware.AdminMiddleware, authMiddleware *middleware.AuthMiddleware, firebase *lib.Firebase, emailClient *lib.EmailClientImpl,
) {
	reactFS, err := fs.Sub(reactApp, "client/dist")
	if err != nil {
		log.Fatal("Failed to create sub filesystem:", err)
	}
	// Serve static files from the embedded filesystem
	router.StaticFS("/assets", http.FS(reactFS))
	// Try to serve favicon and manifest from embedded files
	router.GET("/favicon.svg", func(c *gin.Context) {
		c.FileFromFS("favicon.svg", http.FS(reactFS))
	})
	router.GET("/robots.txt", func(c *gin.Context) {
		c.FileFromFS("robots.txt", http.FS(reactFS))
	})
	router.GET("/sitemap.xml", func(c *gin.Context) {
		c.FileFromFS("sitemap.xml", http.FS(reactFS))
	})
	router.GET("/default.jpg", func(c *gin.Context) {
		c.FileFromFS("default.jpg", http.FS(reactFS))
	})
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Farmec Ireland's API Service.",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "API route not found",
			})
			return
		}
		c.Header("Content-Type", "text/html")
		c.FileFromFS("index.html", http.FS(reactFS))
	})
	// Supplier Module Resources
	routes.InitParts(router, database, s3Client, adminMiddleware)
	routes.InitVideos(router, database, secrets, adminMiddleware)
	routes.InitProduct(router, database, s3Client, adminMiddleware)
	routes.InitMachines(router, database, s3Client, adminMiddleware)
	routes.InitSuppliers(router, database, s3Client, adminMiddleware)
	// About Module Resources
	routes.InitTerms(router, database, adminMiddleware)
	routes.InitPrivacy(router, database, adminMiddleware)
	routes.InitTimelines(router, database, adminMiddleware)
	routes.InitializeEmployee(router, database, s3Client, adminMiddleware)
	// Blog Module Resources
	routes.InitExhibitions(router, database, adminMiddleware)
	routes.InitBlogs(router, database, s3Client, adminMiddleware)
	// Misc Resources
	routes.InitWarranty(router, database, authMiddleware, emailClient)
	routes.InitRegistrations(router, database, authMiddleware, emailClient)
	routes.InitLineItems(router, database, s3Client, adminMiddleware)
	routes.InitCarousel(router, database, s3Client, adminMiddleware)
	// Util Resources
	routes.InitContact(router, emailClient)
	routes.InitCheckout(router, database, secrets)
	routes.InitPdfRenderer(router, adminMiddleware)
	routes.InitAuth(router, firebase, adminMiddleware)
}
