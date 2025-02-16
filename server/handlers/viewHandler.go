package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
	"net/http"
)

type ViewHandler struct {
	carouselService services.CarouselService
	authMiddleware  *middleware.AuthMiddlewareImpl
	supplierCache   *middleware.SupplierCache
}

func NewViewHandler(carouselService services.CarouselService, authMiddleware *middleware.AuthMiddlewareImpl, supplierCahce *middleware.SupplierCache) *ViewHandler {
	return &ViewHandler{carouselService, authMiddleware, supplierCahce}
}

func (handler *ViewHandler) HomeView(c *gin.Context) {

}

func (handler *ViewHandler) CarouselAdminView(context *gin.Context) {
	request := context.Request.Context()
	isAdmin := handler.authMiddleware.GetIsAdmin(context)
	suppliers := handler.supplierCache.GetSuppliersFromContext(context)

	carousels, err := handler.carouselService.GetCarousels(request)
	if err != nil {
		log.Printf("Error getting carousels: %v\n", err)
	}

	isError := err != nil
	component := pages.CarouselAdmin(isAdmin, isError, carousels, suppliers)
	if err := component.Render(request, context.Writer); err != nil {
		log.Printf("Error rendering carousels: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering the page"})
		return
	}
	context.Header("Content-Type", "text/html; charset=utf-8")
}
