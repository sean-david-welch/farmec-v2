package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitContact(router *gin.Engine, secrets *lib.Secrets) {
	loginAuth := lib.NewLoginAuth(secrets.EmailUser, secrets.EmailPass)
	service := services.NewContactService(secrets, loginAuth)
	controller := controllers.NewContactController(service)

	ContactRoutes(router, controller)
}

func ContactRoutes(router *gin.Engine, controller *controllers.ContactController) {
	contactGroup := router.Group("/api/contact")

	contactGroup.POST("", controller.SendEmail)
}
