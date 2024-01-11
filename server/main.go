package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/routes"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func main() {
	secrets, err := config.NewSecrets(); if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	db, err := sql.Open("postgres", secrets.DatabaseURL); if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s3Client, err := utils.NewS3Client("eu-west-1", secrets.AwsAccessKey, secrets.AwsSecret); if err != nil {
		log.Fatal("Failed to create S3 client: ", err)
	}

	firebase, err := lib.NewFirebase(secrets); if err != nil {
		log.Fatal("Failed to initialize Firebase: ", err)
	}
	
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery(), cors.Default())

    routes.InitializeSuppliers(router, db, s3Client, firebase)
	routes.InitializeMachines(router, db, s3Client, firebase)
	routes.InitializeProduct(router, db, s3Client, firebase)
	routes.InitializeVideos(router, db, secrets, firebase)
	routes.InitializeParts(router, db, s3Client, firebase)
	routes.InitializeCarousel(router, db, s3Client, firebase)


	router.Run(":8080")
}