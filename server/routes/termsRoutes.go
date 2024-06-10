package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitTerms(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	termsRepository := repository.NewTermsRepository(database)
	service := services.NewTermsService(termsRepository)
	controller := handlers.NewTermsController(service)

	TermsRoutes(router, controller, adminMiddleware)
}

func TermsRoutes(router *gin.Engine, controller *handlers.TermsController, adminMiddleware *middleware.AdminMiddleware) {
	termsGroup := router.Group("/api/terms")

	termsGroup.GET("", controller.GetTerms)

	protected := termsGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", controller.CreateTerm)
		protected.PUT("/:id", controller.UpdateTerm)
		protected.DELETE("/:id", controller.DeleteTerm)
	}
}
