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
	InitializeParts(router, database, s3Client, adminMiddleware)
	InitializeVideos(router, database, secrets, adminMiddleware)
	InitializeProduct(router, database, s3Client, adminMiddleware)
	InitializeMachines(router, database, s3Client, adminMiddleware)
	InitializeSuppliers(router, database, s3Client, adminMiddleware)

	// About Module Resources
	InitializeTerms(router, database, adminMiddleware)
	InitializePrivacy(router, database, adminMiddleware)
	InitializeTimelines(router, database, adminMiddleware)
	InitilizeEmployee(router, database, s3Client, adminMiddleware)

	// Blog Modeule Resources
	InitializeExhibitions(router, database, adminMiddleware)
	InitializeBlogs(router, database, s3Client, adminMiddleware)
	
	// Misc Resources
	InitializeWarranty(router, database, authMiddleware)
	InitializeRegistrations(router, database, authMiddleware)
	InitializeLineItems(router, database, s3Client, adminMiddleware)
	InitializeCarousel(router, database, s3Client, adminMiddleware)

	// Util Resources
	InitializeContact(router, secrets)
	InitializeCheckout(router, database, secrets)
}