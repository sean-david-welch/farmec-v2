package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/store"
)

func InitTerms(router *gin.Engine, database *sql.DB, adminMiddleware *middleware.AdminMiddleware) {
	termsStore := store.NewTermsStore(database)
	service := services.NewTermsService(termsStore)
	handler := handlers.NewTermsHandler(service)

	TermsRoutes(router, handler, adminMiddleware)
}

func TermsRoutes(router *gin.Engine, handler *handlers.TermsHandler, adminMiddleware *middleware.AdminMiddleware) {
	termsGroup := router.Group("/api/terms")

	termsGroup.GET("", handler.GetTerms)

	protected := termsGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", handler.CreateTerm)
		protected.PUT("/:id", handler.UpdateTerm)
		protected.DELETE("/:id", handler.DeleteTerm)
	}
}
