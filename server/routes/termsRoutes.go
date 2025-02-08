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
	repo := repository.NewTermsRepo(database)
	service := services.NewTermsService(repo)
	handler := handlers.NewTermsHandler(service)

	TermsRoutes(router, handler, adminMiddleware)
}

func TermsRoutes(router *gin.Engine, handler *handlers.TermsHandler, adminMiddleware *middleware.AdminMiddleware) {
	termsGroup := router.Group("/api/terms")

	termsGroup.GET("", handler.GetTerms)

	protected := termsGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.POST("", handler.CreateTerm)
		protected.PUT("/:id", handler.UpdateTerm)
		protected.DELETE("/:id", handler.DeleteTerm)
	}
}
