package main

import (
	"database/sql"
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/routes"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed public/**/*
var embeddedFiles embed.FS

func main() {
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	secrets, err := lib.NewSecrets()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	var database *sql.DB
	var connectionString string

	if env == "production" {
		connectionString = "/home/seanwelch/server/bin/database/database.db"
	} else {
		connectionString = "./database/database.db"
	}

	database, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	s3Client, err := lib.NewS3Client("eu-west-1", secrets.AwsAccessKey, secrets.AwsSecret)
	if err != nil {
		log.Fatal("Failed to create S3 client: ", err)
	}

	firebase, err := lib.NewFirebase(secrets)
	if err != nil {
		log.Fatal("Failed to initialize Firebase: ", err)
	}

	corsConfig := cors.DefaultConfig()
	if env == "production" {
		corsConfig.AllowOrigins = []string{
			"https://farmec.ie",
			"https://www.farmec.ie",
			"https://d2hp5uofb6qy9a.cloudfront.net",
			"https://farmec.ie.s3-website-eu-west-1.amazonaws.com",
		}
	} else {
		corsConfig.AllowOrigins = []string{
			"http://localhost:5173",
		}
	}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Accept", "Origin"}
	corsConfig.AllowCredentials = true

	router := gin.Default()
	emailClient := lib.NewEmailClient(secrets)
	adminMiddleware := middleware.NewAuthMiddleware(firebase)
	supplierMiddleware := middleware.NewSupplierCache(2 * time.Hour)

	router.Use(gin.Logger(), gin.Recovery(), cors.New(corsConfig))
	routes.InitRoutes(router, database, secrets, s3Client, embeddedFiles, firebase, emailClient, adminMiddleware, supplierMiddleware)

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		router.StaticFS("/public", http.FS(embeddedFiles))
		err := router.Run("0.0.0.0:" + port)
		if err != nil {
			return
		}
	} else {
		gin.SetMode(gin.DebugMode)
		router.Static("/public", "./public")
		err := router.Run("localhost:8000")
		if err != nil {
			return
		}
	}
}
