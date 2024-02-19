package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/routes"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func main() {
	secrets, err := config.NewSecrets()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	database, err := sql.Open("postgres", secrets.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	s3Client, err := utils.NewS3Client("eu-west-1", secrets.AwsAccessKey, secrets.AwsSecret)
	if err != nil {
		log.Fatal("Failed to create S3 client: ", err)
	}

	firebase, err := lib.NewFirebase(secrets)
	if err != nil {
		log.Fatal("Failed to initialize Firebase: ", err)
	}

	adminMiddleware := middleware.NewAdminMiddleware(firebase)
	authMiddleware := middleware.NewAuthMiddleware(firebase)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173", "http://farmec.ie.s3-website-eu-west-1.amazonaws.com",
		"https://d2hp5uofb6qy9a.cloudfront.net", "http://farmec.ie",
		"http://www.farmec.ie", "https://farmec.ie", "https://www.farmec.ie",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "Accept", "Origin"}
	config.AllowCredentials = true

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery(), cors.New(config))

	routes.InitRoutes(router, database, secrets, s3Client, adminMiddleware, authMiddleware, firebase)

	port := os.Getenv("PORT")
	router.Run("0.0.0.0:" + port)
}
