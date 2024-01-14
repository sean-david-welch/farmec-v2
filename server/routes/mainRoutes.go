package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB, secrets *config.Secrets, s3Client *utils.S3Client, adminMiddleware *middleware.AdminMiddleware, authMiddleware *middleware.AuthMiddleware) {
	// Supplier Module Resouces
	InitializeSuppliers(router, db, s3Client, adminMiddleware)
	InitializeMachines(router, db, s3Client, adminMiddleware)
	InitializeProduct(router, db, s3Client, adminMiddleware)
	InitializeParts(router, db, s3Client, adminMiddleware)
	InitializeVideos(router, db, secrets, adminMiddleware)

	// About Module Resources
	InitilizeEmployee(router, db, s3Client, adminMiddleware)
	InitializeTimelines(router, db, adminMiddleware)
	InitializePrivacy(router, db, adminMiddleware)
	InitializeTerms(router, db, adminMiddleware)

	// Blog Modeule Resources
	InitializeBlogs(router, db, s3Client, adminMiddleware)
	InitializeExhibitions(router, db, adminMiddleware)

	// Misc Resources
	InitializeCarousel(router, db, s3Client, adminMiddleware)
	InitializeRegistrations(router, db, authMiddleware)
	InitializeWarranty(router, db, authMiddleware)
}