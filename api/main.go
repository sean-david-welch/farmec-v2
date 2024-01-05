package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/sean-david-welch/Farmec-Astro/api/config"
	"github.com/sean-david-welch/Farmec-Astro/api/controllers"
	"github.com/sean-david-welch/Farmec-Astro/api/routes"
	"github.com/sean-david-welch/Farmec-Astro/api/services"
)

func main() {
	config, err := config.NewConfig(); if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	db, err := sql.Open("postgres", config.DatabaseURL); if err != nil {
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