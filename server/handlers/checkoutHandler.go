package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type CheckoutHandler struct {
	service services.CheckoutService
}

func NewCheckoutHandler(service services.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{service: service}
}

func (handler *CheckoutHandler) CreateCheckoutSession(context *gin.Context) {
	id := context.Param("id")

	sess, err := handler.service.CreateCheckoutSession(id)
	if err != nil {
		log.Printf("error occurred in checkout service: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred when trying to create checkout session"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"session":       sess,
		"clientSecret:": sess.ClientSecret,
	})
}

func (handler *CheckoutHandler) RetrieveCheckoutSession(context *gin.Context) {
	sessionId := context.Query("session_id")

	if sessionId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	sess, err := handler.service.RetrieveCheckoutSession(sessionId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":         sess.Status,
		"customer_email": sess.CustomerDetails.Email,
	})
}
