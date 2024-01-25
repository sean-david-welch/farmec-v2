package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitWarranty(router *gin.Engine, database *sql.DB, authMiddleware *middleware.AuthMiddleware) {
	repository := repository.NewWarrantyRepository(database)
	service := services.NewWarrantyService(repository)
	controller := controllers.NewWarrantyController(service)

	WarrantyRoutes(router, controller, authMiddleware)
}

func WarrantyRoutes(router *gin.Engine, controller *controllers.WarrantyController, authMiddleware *middleware.AuthMiddleware) {
	warrantyGroup := router.Group("/api/warranty")

	warrantyGroup.GET("", controller.GetWarranties)
	warrantyGroup.GET("/:id", controller.GetWarrantyById)

	protected := warrantyGroup.Group("").Use(authMiddleware.Middleware())
	{
		protected.POST("", controller.CreateWarranty)
		protected.PUT("/:id", controller.UpdateWarranty)
		protected.DELETE("/:id", controller.DeleteWarranty)
	}
}
