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

func InitMachines(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	repo := repository.NewMachineRepo(database)
	service := services.NewMachineService(repo, s3Client, "Machines")
	handler := handlers.NewMachineHandler(service)

	MachineRoutes(router, handler, adminMiddleware)
}

func MachineRoutes(router *gin.Engine, handler *handlers.MachineHandler, adminMiddleware *middleware.AdminMiddleware) {
	machineGroup := router.Group("/api/machines")

	machineGroup.GET("/:id", handler.GetMachineById)
	machineGroup.GET("/suppliers/:id", handler.GetMachines)

	protected := machineGroup.Group("").Use(adminMiddleware.RouteMiddleware())
	{
		protected.POST("", handler.CreateMachine)
		protected.PUT("/:id", handler.UpdateMachine)
		protected.DELETE("/:id", handler.DeleteMachine)
	}
}
