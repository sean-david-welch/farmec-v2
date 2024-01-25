package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitContact(router *gin.Engine, secrets *config.Secrets) {
	service := services.NewContactService(secrets)
	controller := controllers.NewContactController(service)

	ContactRoutes(router, controller)
}

func ContactRoutes(router *gin.Engine, controller *controllers.ContactController) {
	contactGroup := router.Group("/api/contact")

	contactGroup.POST("", controller.SendEmail)
}
