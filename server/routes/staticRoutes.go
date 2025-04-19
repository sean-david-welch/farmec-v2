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
	router.GET("/assets/*filepath", func(c *gin.Context) {
		path := c.Param("filepath")
		if path == "/" {
			c.Status(http.StatusNotFound)
			return
		}
		path = strings.TrimPrefix(path, "/")

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

		assetPath := "assets/" + path
		log.Printf("Looking for file at: %s", assetPath)

		file, err := reactApp.Open(assetPath)
		if err != nil {
			log.Printf("Error opening asset: %v", err)
			c.Status(http.StatusNotFound)
			return
		}
		defer func(file fs.File) {
			err := file.Close()
			if err != nil {
				log.Printf("Error closing asset: %v", err)
			}
		}(file)

		content, err := io.ReadAll(file)
		if err != nil {
			log.Printf("Error reading asset: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		_, err = c.Writer.Write(content)
		if err != nil {
			return
		}
	})

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

func serveStaticFile(filename string, file fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := file.Open(filename)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		defer func(file fs.File) {
			err := file.Close()
			if err != nil {
				log.Printf("Error closing asset: %v", err)
			}
		}(file)

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
