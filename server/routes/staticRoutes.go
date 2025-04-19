package routes

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

func staticRoutes(router *gin.Engine, reactApp fs.FS) {
	router.GET("/", func(c *gin.Context) {
		indexHTML, err := fs.ReadFile(reactApp, "index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read index.html")
			return
		}

		c.Header("Content-Type", "text/html")
		c.Status(http.StatusOK)
		_, err = c.Writer.Write(indexHTML)
		if err != nil {
			return
		}
		c.Abort()
	})
	// Custom static file server with correct MIME types
	router.GET("/assets/*filepath", func(c *gin.Context) {
		// The path parameter includes the leading slash, so we need to adjust
		path := c.Param("filepath")

		// Skip requests to just "/assets/" without a specific file
		if path == "/" {
			c.Status(http.StatusNotFound)
			return
		}

		// Remove the leading slash for fs.Open
		path = strings.TrimPrefix(path, "/")

		log.Printf("Requesting asset: %s", path)

		// Set correct MIME type based on file extension
		if strings.HasSuffix(path, ".js") {
			c.Header("Content-Type", "application/javascript")
		} else if strings.HasSuffix(path, ".css") {
			c.Header("Content-Type", "text/css")
		} else if strings.HasSuffix(path, ".svg") {
			c.Header("Content-Type", "image/svg+xml")
		} else if strings.HasSuffix(path, ".png") {
			c.Header("Content-Type", "image/png")
		} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
			c.Header("Content-Type", "image/jpeg")
		}

		// Construct the path correctly for the embedded filesystem
		assetPath := "assets/" + path
		log.Printf("Looking for file at: %s", assetPath)

		file, err := reactApp.Open(assetPath)
		if err != nil {
			log.Printf("Error opening asset: %v", err)
			c.Status(http.StatusNotFound)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error reading asset: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Writer.Write(content)
	})

	// Serve other static files directly with correct MIME types
	router.GET("/favicon.svg", func(c *gin.Context) {
		c.Header("Content-Type", "image/svg+xml")
		serveStaticFile("favicon.svg", reactApp)(c)
	})

	router.GET("/robots.txt", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		serveStaticFile("robots.txt", reactApp)(c)
	})

	router.GET("/sitemap.xml", func(c *gin.Context) {
		c.Header("Content-Type", "application/xml")
		serveStaticFile("sitemap.xml", reactApp)(c)
	})

	router.GET("/default.jpg", func(c *gin.Context) {
		c.Header("Content-Type", "image/jpeg")
		serveStaticFile("default.jpg", reactApp)(c)
	})

	// NoRoute handler for SPA routing - serve index.html for all unmatched routes except API
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.HasPrefix(path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "API route not found",
			})
			return
		}

		indexHTML, err := fs.ReadFile(reactApp, "index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read index.html")
			return
		}

		c.Header("Content-Type", "text/html")
		c.Status(http.StatusOK)
		_, err = c.Writer.Write(indexHTML)
		if err != nil {
			return
		}
	})

}

func serveStaticFile(filename string, fs fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := fs.Open(filename)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		_, err = c.Writer.Write(content)
		if err != nil {
			return
		}
	}
}
