package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/routes"
)

func main() {
	secrets, err := config.NewSecrets(); if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	db, err := sql.Open("postgres", secrets.DatabaseURL); if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	router := gin.Default()

	router.Use(gin.Logger(), gin.Recovery(), cors.Default())

    routes.InitializeSupplier(router, db)


	router.Run(":8080")
}