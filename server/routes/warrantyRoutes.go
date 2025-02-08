package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitWarranty(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AuthMiddlewareImpl, smtp *lib.SMTPClientImpl) {
	repo := repository.NewWarrantyRepo(database)
	service := services.NewWarrantyService(repo, *smtp)
	handler := handlers.NewWarrantyHandler(service)

	WarrantyRoutes(router, handler, adminMiddleware)
}

func WarrantyRoutes(router *gin.Engine, handler *handlers.WarrantyHandler, adminMiddleware *middleware.AuthMiddlewareImpl) {
	warrantyGroup := router.Group("/api/warranty")

	warrantyGroup.GET("", handler.GetWarranties)
	warrantyGroup.GET("/:id", handler.GetWarrantyById)
	warrantyGroup.POST("", handler.CreateWarranty)

	protected := warrantyGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.PUT("/:id", handler.UpdateWarranty)
		protected.DELETE("/:id", handler.DeleteWarranty)
	}
}
