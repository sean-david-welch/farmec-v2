package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitRoutes(router *gin.Engine, database *sql.DB, secrets *config.Secrets, s3Client utils.S3Client, adminMiddleware *middleware.AdminMiddleware, authMiddleware *middleware.AuthMiddleware, firebase *lib.Firebase) {
	// Supplier Module Resouces
	InitParts(router, database, s3Client, adminMiddleware)
	InitVideos(router, database, secrets, adminMiddleware)
	InitProduct(router, database, s3Client, adminMiddleware)
	InitMachines(router, database, s3Client, adminMiddleware)
	InitSuppliers(router, database, s3Client, adminMiddleware)

	// About Module Resources
	InitTerms(router, database, adminMiddleware)
	InitPrivacy(router, database, adminMiddleware)
	InitTimelines(router, database, adminMiddleware)
	InitilizeEmployee(router, database, s3Client, adminMiddleware)

	// Blog Modeule Resources
	InitExhibitions(router, database, adminMiddleware)
	InitBlogs(router, database, s3Client, adminMiddleware)

	// Misc Resources
	InitWarranty(router, database, authMiddleware)
	InitRegistrations(router, database, authMiddleware)
	InitLineItems(router, database, s3Client, adminMiddleware)
	InitCarousel(router, database, s3Client, adminMiddleware)

	// Util Resources
	InitContact(router, secrets)
	InitCheckout(router, database, secrets)
	InitAuth(router, firebase, adminMiddleware)
}