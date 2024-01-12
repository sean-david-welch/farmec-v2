package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitializeTerms(router *gin.Engine, db *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	repository := repository.NewTermsRepository(db)
	service := services.NewTermsService(repository)
	controller := controllers.NewTermsController(service)

	TermsRoutes(router, controller, adminMiddleware)
}

func TermsRoutes(router *gin.Engine, controller *controllers.TermsController, adminMiddleware *middleware.AdminMiddleware) {
	termsGroup := router.Group("/api/terms")

	termsGroup.GET("", controller.GetTerms)

	protected := termsGroup.Group("").Use(adminMiddleware.Middleware()); {
		protected.POST("", controller.CreateTerm)
		protected.PUT("/:id", controller.UpdateTerm)
		protected.DELETE("/:id", controller.DeleteTerm)
	}
}