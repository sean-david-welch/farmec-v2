package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type TermsController struct {
	service services.TermsService
}

func NewTermsController(service services.TermsService) *TermsController {
	return &TermsController{service: service}
}

func (controller *TermsController) GetTerms(context *gin.Context) {
	terms, err := controller.service.GetTerms()
	if err != nil {
		log.Printf("error getting terms: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting term"})
		return
	}

	context.JSON(http.StatusOK, terms)
}

func (controller *TermsController) CreateTerm(context *gin.Context) {
	var term types.Terms

	if err := context.ShouldBindJSON(&term); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := controller.service.CreateTerm(&term); err != nil {
		log.Printf("error while creating term: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating term", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, term)
}

func (controller *TermsController) UpdateTerm(context *gin.Context) {
	id := context.Param("id")
	var term types.Terms

	if err := context.ShouldBindJSON(&term); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := controller.service.UpdateTerm(id, &term); err != nil {
		log.Printf("error while updating term: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred while updating term", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, term)
}

func (controller *TermsController) DeleteTerm(context *gin.Context) {
	id := context.Param("id")

	if err := controller.service.DeleteTerm(id); err != nil {
		log.Printf("Error deleting term: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting term", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "term deleted successfully", "id": id})
}
