package handlers

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type Handlers struct {
	SupplierHandler *SupplierHandler
	CarouselHandler *CarouselHandler
	ViewHandler     *ViewHandler
	// Add other handlers here
}

func NewHandlers(
	db *sql.DB,
	s3Client lib.S3Client,
	authMiddleware *middleware.AuthMiddlewareImpl,
	supplierCache *middleware.SupplierCache,
) *Handlers {
	// Initialize repositories
	supplierRepo := repository.NewSupplierRepo(db)
	carouselRepo := repository.NewCarouselRepo(db)

	// Initialize services
	supplierService := services.NewSupplierService(supplierRepo, s3Client, "Suppliers")
	carouselService := services.NewCarouselService(carouselRepo, s3Client, "Carousels")

	// Initialize handlers
	return &Handlers{
		SupplierHandler: NewSupplierHandler(supplierService),
		CarouselHandler: NewCarouselHandler(carouselService, authMiddleware, supplierCache),
		ViewHandler:     NewViewHandler(carouselService, authMiddleware, supplierCache),
	}
}
