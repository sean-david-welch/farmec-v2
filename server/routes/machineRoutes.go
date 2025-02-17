package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
)

func MachineRoutes(router *gin.Engine, handler *handlers.MachineHandler, authMiddleware *middleware.AuthMiddlewareImpl) {
	machineGroup := router.Group("/api/machines")

	machineGroup.GET("/:id", handler.GetMachineById)
	machineGroup.GET("/suppliers/:id", handler.GetMachines)

	protected := machineGroup.Group("").Use(authMiddleware.AdminRouteMiddleware())
	{
		protected.POST("", handler.CreateMachine)
		protected.PUT("/:id", handler.UpdateMachine)
		protected.DELETE("/:id", handler.DeleteMachine)
	}
}
