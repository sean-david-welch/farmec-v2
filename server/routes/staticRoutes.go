package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

func staticRoutes(router *gin.Engine, queries *db.Queries, reactApp fs.FS) {
	router.GET("/", func(context *gin.Context) {
		indexHTML, err := fs.ReadFile(reactApp, "index.html")
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to read index.html")
			return
		}

		context.Header("Content-Type", "text/html")
		context.Status(http.StatusOK)
		_, err = context.Writer.Write(indexHTML)
		if err != nil {
			return
		}
		context.Abort()
	})
	router.GET("/assets/*filepath", func(context *gin.Context) {
		path := context.Param("filepath")
		if path == "/" {
			context.Status(http.StatusNotFound)
			return
		}
		path = strings.TrimPrefix(path, "/")

		if strings.HasSuffix(path, ".js") {
			context.Header("Content-Type", "application/javascript")
		} else if strings.HasSuffix(path, ".css") {
			context.Header("Content-Type", "text/css")
		} else if strings.HasSuffix(path, ".svg") {
			context.Header("Content-Type", "image/svg+xml")
		} else if strings.HasSuffix(path, ".png") {
			context.Header("Content-Type", "image/png")
		} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
			context.Header("Content-Type", "image/jpeg")
		}

		assetPath := "assets/" + path
		log.Printf("Looking for file at: %s", assetPath)

		file, err := reactApp.Open(assetPath)
		if err != nil {
			log.Printf("Error opening asset: %v", err)
			context.Status(http.StatusNotFound)
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
			context.Status(http.StatusInternalServerError)
			return
		}

		_, err = context.Writer.Write(content)
		if err != nil {
			return
		}
	})

	router.GET("/favicon.svg", func(context *gin.Context) {
		context.Header("Content-Type", "image/svg+xml")
		serveStaticFile("favicon.svg", reactApp)(context)
	})

	router.GET("/robots.txt", func(context *gin.Context) {
		context.Header("Content-Type", "text/plain")
		serveStaticFile("robots.txt", reactApp)(context)
	})

	router.GET("/sitemap.xml", func(context *gin.Context) {
		context.Header("Content-Type", "application/xml")
		serveStaticFile("sitemap.xml", reactApp)(context)
	})

	router.GET("/default.jpg", func(context *gin.Context) {
		context.Header("Content-Type", "image/jpeg")
		serveStaticFile("default.jpg", reactApp)(context)
	})

	router.GET("/api/sitemap-data", func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		sitemapData := lib.SitemapData(context, queries)
		context.JSON(http.StatusOK, sitemapData)
	})

	router.NoRoute(func(context *gin.Context) {
		path := context.Request.URL.Path

		if strings.HasPrefix(path, "/api") {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "API route not found",
			})
			return
		}

		indexHTML, err := fs.ReadFile(reactApp, "index.html")
		if err != nil {
			context.String(http.StatusInternalServerError, "Failed to read index.html")
			return
		}

		context.Header("Content-Type", "text/html")
		context.Status(http.StatusOK)
		_, err = context.Writer.Write(indexHTML)
		if err != nil {
			return
		}
	})

}

func serveStaticFile(filename string, file fs.FS) gin.HandlerFunc {
	return func(context *gin.Context) {
		file, err := file.Open(filename)
		if err != nil {
			context.Status(http.StatusNotFound)
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
			context.Status(http.StatusInternalServerError)
			return
		}

		_, err = context.Writer.Write(content)
		if err != nil {
			return
		}
	}
}
