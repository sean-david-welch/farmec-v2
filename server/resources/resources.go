package resources

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type Resources struct {
	SupplierService services.SupplierService
	SupplierHandler *handlers.SupplierHandler
	CarouselHandler *handlers.CarouselHandler
	ViewHandler     *handlers.ViewHandler
	// Add other handlers here
}

func NewResources(
	db *sql.DB, s3Client lib.S3Client, authMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache,
) *Resources {
	// Initialize repositories
	supplierRepo := repository.NewSupplierRepo(db)
	carouselRepo := repository.NewCarouselRepo(db)

	// Initialize services
	supplierService := services.NewSupplierService(supplierRepo, s3Client, "Suppliers")
	carouselService := services.NewCarouselService(carouselRepo, s3Client, "Carousels")

	// Initialize handlers
	return &Resources{
		SupplierService: supplierService,
		SupplierHandler: handlers.NewSupplierHandler(supplierService),
		CarouselHandler: handlers.NewCarouselHandler(carouselService, authMiddleware, supplierCache),
		ViewHandler:     handlers.NewViewHandler(carouselService, authMiddleware, supplierCache),
	}
}
