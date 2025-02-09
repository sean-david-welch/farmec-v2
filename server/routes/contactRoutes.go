package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

func InitContact(router *gin.Engine, smtp *lib.EmailClientImpl) {
	service := services.NewContactService(smtp)
	handler := handlers.NewContactHandler(service)

	router.POST("/api/contact", handler.SendEmail)
}
