package routes

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
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

	router.StaticFS("/assets", http.FS(reactApp))
	// Custom static file server with correct MIME types
	router.GET("/assets/*filepath", func(c *gin.Context) {
		path := c.Param("filepath")

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

		// Read and serve the file directly
		file, err := reactApp.Open("assets" + path)
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
