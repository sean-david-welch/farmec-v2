package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeMachines(router *gin.Engine, db *sql.DB, s3Client *utils.S3Client, adminMiddleware *middleware.AdminMiddleware) {
	machineRepository := repository.NewMachineRepository(db)
	machineService := services.NewMachineService(machineRepository, s3Client, "machines")
	machineController := controllers.NewMachineController(machineService)

	MachineRoutes(router, machineController, adminMiddleware)
}

func MachineRoutes(router *gin.Engine, machineController *controllers.MachineController, adminMiddleware *middleware.AdminMiddleware) {
	machineGroup := router.Group("/api/machines")

	machineGroup.GET("/:id", machineController.GetMachines)
	
	protected := machineGroup.Group("").Use(adminMiddleware.Middleware()); {
		protected.POST("", machineController.CreateMachine)
		protected.PUT("/:id", machineController.UpdateMachine)
		protected.DELETE("/:id", machineController.DeleteMachine)
	}
}