package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type ViewHandler struct {
	carouselService services.CarouselService
	authMiddleware  *middleware.AuthMiddlewareImpl
	supplierCache   *middleware.SupplierCache
}

func NewViewHandler(carouselService services.CarouselService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCahce *middleware.SupplierCache) *ViewHandler {
	return &ViewHandler{carouselService, authMiddleware, supplierCahce}
}
