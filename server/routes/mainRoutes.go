package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"io/fs"
)

func InitRoutes(
	reactApp fs.FS, router *gin.Engine, database *sql.DB, secrets *lib.Secrets, s3Client lib.S3Client,
	adminMiddleware *middleware.AdminMiddleware, authMiddleware *middleware.AuthMiddleware, firebase *lib.Firebase, emailClient *lib.EmailClientImpl,
) {
	staticRoutes(router, reactApp)
	// Supplier Module Resources
	InitParts(router, database, s3Client, adminMiddleware)
	InitVideos(router, database, secrets, adminMiddleware)
	InitProduct(router, database, s3Client, adminMiddleware)
	InitMachines(router, database, s3Client, adminMiddleware)
	InitSuppliers(router, database, s3Client, adminMiddleware)
	// About Module Resources
	InitTerms(router, database, adminMiddleware)
	InitPrivacy(router, database, adminMiddleware)
	InitTimelines(router, database, adminMiddleware)
	InitializeEmployee(router, database, s3Client, adminMiddleware)
	// Blog Module Resources
	InitExhibitions(router, database, adminMiddleware)
	InitBlogs(router, database, s3Client, adminMiddleware)
	// Misc Resources
	InitWarranty(router, database, authMiddleware, emailClient)
	InitRegistrations(router, database, authMiddleware, emailClient)
	InitLineItems(router, database, s3Client, adminMiddleware)
	InitCarousel(router, database, s3Client, adminMiddleware)
	// Util Resources
	InitContact(router, emailClient)
	InitCheckout(router, database, secrets)
	InitPdfRenderer(router, adminMiddleware)
	InitAuth(router, firebase, adminMiddleware)
}
