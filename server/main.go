package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/routes"
)

func main() {
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	secrets, err := lib.NewSecrets()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	database, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	s3Client, err := lib.NewS3Client("eu-west-1", secrets.AwsAccessKey, secrets.AwsSecret)
	if err != nil {
		log.Fatal("Failed to create S3 client: ", err)
	}

	firebase, err := lib.NewFirebase(secrets)
	if err != nil {
		log.Fatal("Failed to initialize Firebase: ", err)
	}

	adminMiddleware := middleware.NewAdminMiddleware(firebase)
	authMiddleware := middleware.NewAuthMiddleware(firebase)

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
	router.Use(gin.Logger(), gin.Recovery(), cors.New(corsConfig))

	routes.InitRoutes(router, database, secrets, s3Client, adminMiddleware, authMiddleware, firebase)

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		err := router.Run("0.0.0.0:" + port)
		if err != nil {
			return
		}
	} else {
		gin.SetMode(gin.DebugMode)
		err := router.Run("127.0.0.1:8000")
		if err != nil {
			return
		}
	}
}
