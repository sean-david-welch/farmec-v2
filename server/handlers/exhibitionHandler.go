package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/views"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type ExhibitionHandler struct {
	service services.ExhibitionService
}

func NewExhibitionHandler(service services.ExhibitionService) *ExhibitionHandler {
	return &ExhibitionHandler{service: service}
}

func (handler *ExhibitionHandler) GetExhibitions(c *gin.Context) {
	ctx := c.Request.Context()
	exhibitions, err := handler.service.GetExhibitions(ctx)
	if err != nil {
		log.Printf("error getting exhibitions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting exhibitions"})
		return
	}

	// Convert to []types.Exhibition if necessary
	exhibitionSlice := make([]types.Exhibition, len(exhibitions))
	for i, e := range exhibitions {
		exhibitionSlice[i] = e
	}

	// Render the Templ component
	component := views.Exhibitions(true, exhibitionSlice, false, false)
	err = component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		log.Printf("error rendering template: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while rendering the page"})
		return
	}

	// Set the content type to HTML
	c.Header("Content-Type", "text/html")
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
