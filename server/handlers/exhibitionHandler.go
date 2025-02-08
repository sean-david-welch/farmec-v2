package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/middleware"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views/pages"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type ExhibitionHandler struct {
	service         services.ExhibitionService
	adminMiddleware *middleware.AuthMiddlewareImpl
	supplierCache   *middleware.SupplierCache
}

func NewExhibitionHandler(service services.ExhibitionService, adminMiddleware *middleware.AuthMiddlewareImpl, supplierCache *middleware.SupplierCache) *ExhibitionHandler {
	return &ExhibitionHandler{service: service, adminMiddleware: adminMiddleware, supplierCache: supplierCache}
}

func (handler *ExhibitionHandler) ExhibitionsView(context *gin.Context) {
	ctx := context.Request.Context()
	isAdmin := handler.adminMiddleware.GetIsAdmin(context)
	suppliers := middleware.GetSuppliersFromContext(context)

	exhibitions, err := handler.service.GetExhibitions(ctx)
	if err != nil {
		log.Printf("error getting exhibitions: %v", err)
	}

	isError := err != nil
	component := pages.Exhibitions(isAdmin, isError, exhibitions, suppliers)
	if err := component.Render(ctx, context.Writer); err != nil {
		log.Printf("error rendering template: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering the page"})
		return
	}
	context.Header("Content-Type", "text/html")
}

func (handler *ExhibitionHandler) GetExhibitions(context *gin.Context) {
	ctx := context.Request.Context()
	exhibitions, err := handler.service.GetExhibitions(ctx)
	if err != nil {
		log.Printf("error getting exhibitions: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting exhibitions"})
		return
	}

	context.JSON(http.StatusOK, exhibitions)
}

func (handler *ExhibitionHandler) CreateExhibition(context *gin.Context) {
	ctx := context.Request.Context()
	var exhibition types.Exhibition

	if err := context.ShouldBindJSON(&exhibition); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbExhibition := lib.DeserializeExhibition(exhibition)
	if err := handler.service.CreateExhibition(ctx, &dbExhibition); err != nil {
		log.Printf("error occurred while creating exhibition: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating exhibition", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, exhibition)
}

func (handler *ExhibitionHandler) UpdateExhibition(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	var exhibition types.Exhibition
	if err := context.ShouldBindJSON(&exhibition); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	dbExhibition := lib.DeserializeExhibition(exhibition)
	if err := handler.service.UpdateExhibition(ctx, id, &dbExhibition); err != nil {
		log.Printf("error occurred while updating exhibition: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating exhibition", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, exhibition)
}

func (handler *ExhibitionHandler) DeleteExhibition(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteExhibition(ctx, id); err != nil {
		log.Printf("error occurred while deleting exhibition: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting exhibition"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "exhibition deleted successfully", "id": id})
}
