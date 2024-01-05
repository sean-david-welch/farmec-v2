package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	secrets "github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/routes"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func main() {
	secrets, err := secrets.NewSecrets(); if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	db, err := sql.Open("postgres", secrets.DatabaseURL); if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	router := gin.Default()

	router.Use(gin.Logger(), gin.Recovery(), cors.Default())

	supplierService := services.NewSupplierService(db)
	supplierController := controllers.NewSuppliersContoller(supplierService)
	routes.SupplierRoutes(router, supplierController)


	router.Run(":8080")
}