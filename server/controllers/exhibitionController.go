package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type ExhibitionController struct {
	service *services.ExhibitionService
}

func NewExhibitionController(service *services.ExhibitionService) *ExhibitionController {
	return &ExhibitionController{service: service}
}

func(controller *ExhibitionController) GetExhibitions(ctx *gin.Context) {
	exhibitions, err := controller.service.GetExhibitions(); if err != nil {
		log.Printf("error getting exhibitions: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting exhibitions"})
		return
	}

	ctx.JSON(http.StatusOK, exhibitions)
}

func(controller *ExhibitionController) CreateExhibition(ctx *gin.Context) {
	var exhibition *models.Exhibition

	if err := ctx.ShouldBindJSON(&exhibition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := controller.service.CreateExhibition(exhibition); err != nil {
		log.Printf("error occurred while creating exhibition: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating exhibition", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, exhibition)
}

func(controller *ExhibitionController) UpdateExhibition(ctx *gin.Context) {
	id := ctx.Param("id")
	
	var exhibition *models.Exhibition

	if err := ctx.ShouldBindJSON(&exhibition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	if err := controller.service.UpdateExhibition(id, exhibition); err != nil {
		log.Printf("error occurred while updating exhibition: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating exhibition", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, exhibition)
}

func(controller *ExhibitionController) DeleteExhibition(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := controller.service.DeleteExhibition(id); err != nil {
		log.Printf("error occurred while deleting exhibition: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting exhibition"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "exhibition deleted successfully", "id": id})
}

