package routes

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/stores"
)

func InitMachines(router *gin.Engine, database *sql.DB, s3Client lib.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	machineStore := stores.NewMachineStore(database)
	machineService := services.NewMachineService(machineStore, s3Client, "Machines")
	machineHandler := handlers.NewMachineHandler(machineService)

	MachineRoutes(router, machineHandler, adminMiddleware)
}

func MachineRoutes(router *gin.Engine, machineHandler *handlers.MachineHandler, adminMiddleware *middleware.AdminMiddleware) {
	machineGroup := router.Group("/api/machines")

	machineGroup.GET("/:id", machineHandler.GetMachineById)
	machineGroup.GET("/suppliers/:id", machineHandler.GetMachines)

	protected := machineGroup.Group("").Use(adminMiddleware.Middleware())
	{
		protected.POST("", machineHandler.CreateMachine)
		protected.PUT("/:id", machineHandler.UpdateMachine)
		protected.DELETE("/:id", machineHandler.DeleteMachine)
	}
}
