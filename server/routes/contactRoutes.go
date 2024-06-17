package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitContact(router *gin.Engine, smtp lib.SMTPClient) {
	service := services.NewContactService(smtp)
	handler := handlers.NewContactHandler(service)

	ContactRoutes(router, handler)
}

func ContactRoutes(router *gin.Engine, handler *handlers.ContactHandler) {
	contactGroup := router.Group("/api/contact")

	contactGroup.POST("", handler.SendEmail)
}
