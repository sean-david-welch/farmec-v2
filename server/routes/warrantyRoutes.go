package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitWarranty(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddleware) {
	warrantyRepository := repository.NewWarrantyRepository(database)
	service := services.NewWarrantyService(warrantyRepository)
	controller := handlers.NewWarrantyController(service)

	WarrantyRoutes(router, controller, authMiddleware)
}

func WarrantyRoutes(router *gin.Engine, controller *handlers.WarrantyController, authMiddleware *middleware.AuthMiddleware) {
	warrantyGroup := router.Group("/api/warranty")

	warrantyGroup.GET("", controller.GetWarranties)
	warrantyGroup.GET("/:id", controller.GetWarrantyById)
	warrantyGroup.POST("", controller.CreateWarranty)

	protected := warrantyGroup.Group("").Use(authMiddleware.Middleware())
	{
		protected.PUT("/:id", controller.UpdateWarranty)
		protected.DELETE("/:id", controller.DeleteWarranty)
	}
}
