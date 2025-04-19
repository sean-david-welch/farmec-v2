package main

import (
	"database/sql"
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/routes"
	"io/fs"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

//go:embed client/dist
var reactApp embed.FS

func main() {
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")
	reactFS, err := fs.Sub(reactApp, "client/dist")
	if err != nil {
		log.Fatal("Failed to create sub filesystem:", err)
	}

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
	authMiddleware := middleware.NewAuthMiddleware(firebase)
	adminMiddleware := middleware.NewAdminMiddleware(firebase)
	prerenderMiddleware := middleware.NewPrerenderMiddleware(secrets.PrerenderToken)
	router.Use(gin.Logger(), gin.Recovery(), cors.New(corsConfig), prerenderMiddleware.Middleware())
	routes.InitRoutes(reactFS, router, database, secrets, s3Client, adminMiddleware, authMiddleware, firebase, emailClient)

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		err := router.Run("0.0.0.0:" + port)
		if err != nil {
			return
		}
	} else {
		gin.SetMode(gin.DebugMode)
		err := router.Run("localhost:8000")
		if err != nil {
			return
		}
	}
}
