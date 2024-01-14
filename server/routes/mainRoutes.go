package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeRoutes(router *gin.Engine, database *sql.DB, secrets *config.Secrets, s3Client *utils.S3Client, adminMiddleware *middleware.AdminMiddleware, authMiddleware *middleware.AuthMiddleware) {
	// Supplier Module Resouces
	InitializeSuppliers(router, database, s3Client, adminMiddleware)
	InitializeMachines(router, database, s3Client, adminMiddleware)
	InitializeProduct(router, database, s3Client, adminMiddleware)
	InitializeParts(router, database, s3Client, adminMiddleware)
	InitializeVideos(router, database, secrets, adminMiddleware)

	// About Module Resources
	InitilizeEmployee(router, database, s3Client, adminMiddleware)
	InitializeTimelines(router, database, adminMiddleware)
	InitializePrivacy(router, database, adminMiddleware)
	InitializeTerms(router, database, adminMiddleware)

	// Blog Modeule Resources
	InitializeBlogs(router, database, s3Client, adminMiddleware)
	InitializeExhibitions(router, database, adminMiddleware)

	// Misc Resources
	InitializeCarousel(router, database, s3Client, adminMiddleware)
	InitializeRegistrations(router, database, authMiddleware)
	InitializeWarranty(router, database, authMiddleware)
	InitializeLineItems(router, database, adminMiddleware)
}