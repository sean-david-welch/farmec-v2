package main

import (
	"database/sql"
	"log"
	"os"

	"githib.com/sean-david-welch/Farmec-Astro/api/controllers"
	"githib.com/sean-david-welch/Farmec-Astro/api/routes"
	"githib.com/sean-david-welch/Farmec-Astro/api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
        log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	supplierService := services.NewSupplierService(db)
	supplierController := controllers.NewSuppliersContoller(supplierService)
	routes.SupplierRoutes(router, supplierController)

	router.Run(":8080")
}