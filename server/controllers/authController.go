package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (controller *AuthController) Logout(context *gin.Context) {
	context.SetCookie("session", "", -1, "/", "", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (controller *AuthController) Login(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")

	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "an error occurred, could not find auth header"})
		return
	}

	idToken := strings.TrimPrefix(authHeader, "Bearer ")

	sessionCookie, err := controller.service.Login(context.Request.Context(), idToken)
	if err != nil {
		log.Printf("error with validating token: %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	context.SetCookie("session", sessionCookie, 72*3600, "/", "", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (controller *AuthController) Register(context *gin.Context) {
	var userData types.UserData

	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	err := controller.service.Register(context, &userData)
	if err != nil {
		log.Printf("error with creating user in firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating user in firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user successfully created in firebase"})
}