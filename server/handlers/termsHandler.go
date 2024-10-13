package handlers

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsHandler struct {
	service services.TermsService
}

func NewTermsHandler(service services.TermsService) *TermsHandler {
	return &TermsHandler{service: service}
}

func (handler *TermsHandler) GetTerms(context *gin.Context) {
	ctx := context.Request.Context()
	terms, err := handler.service.GetTerms(ctx)
	if err != nil {
		log.Printf("error getting terms: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting term"})
		return
	}

	context.JSON(http.StatusOK, terms)
}

func (handler *TermsHandler) CreateTerm(context *gin.Context) {
	ctx := context.Request.Context()
	var term types.Terms
	dbTerm := lib.DeserializeTerm(term)

	if err := context.ShouldBindJSON(&term); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := handler.service.CreateTerm(ctx, &dbTerm); err != nil {
		log.Printf("error while creating term: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating term", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, term)
}

func (handler *TermsHandler) UpdateTerm(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")
	var term types.Terms
	dbTerm := lib.DeserializeTerm(term)

	if err := context.ShouldBindJSON(&term); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := handler.service.UpdateTerm(ctx, id, &dbTerm); err != nil {
		log.Printf("error while updating term: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating term", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, term)
}

func (handler *TermsHandler) DeleteTerm(context *gin.Context) {
	ctx := context.Request.Context()
	id := context.Param("id")

	if err := handler.service.DeleteTerm(ctx, id); err != nil {
		log.Printf("Error deleting term: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting term", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "term deleted successfully", "id": id})
}
