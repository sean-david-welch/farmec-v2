package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Logout(context *gin.Context) {
	context.SetCookie("access_token", "", -1, "/", "", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (handler *AuthHandler) Login(context *gin.Context) {
	env := os.Getenv("ENV")

	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "an error occurred, could not find auth header"})
		return
	}

	idToken := strings.TrimPrefix(authHeader, "Bearer ")

	sessionCookie, err := handler.service.Login(context.Request.Context(), idToken)
	if err != nil {
		log.Printf("error with validating token: %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	if env == "production" {
		context.SetCookie("access_token", sessionCookie, 72*3600, "/", "farmec.ie", true, true)
	} else {
		context.SetCookie("access_token", sessionCookie, 72*3600, "/", "localhost", false, true)
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (handler *AuthHandler) GetUsers(context *gin.Context) {
	users, err := handler.service.GetUsers(context)
	if err != nil {
		log.Printf("error fetching users from firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching users from firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var userData types.UserData

	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	err := handler.service.Register(context, &userData)
	if err != nil {
		log.Printf("error with creating user in firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating user in firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user successfully created in firebase"})
}

func (handler *AuthHandler) UpdateUser(context *gin.Context) {
	uid := context.Param("uid")
	var userData types.UserData

	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating user in firebase", "details": err.Error()})
		return
	}

	if err := handler.service.UpdateUser(uid, &userData, context); err != nil {
		log.Printf("error while updating user in firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating user in firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user successfully updated in firebase"})
}

func (handler *AuthHandler) DeleteUser(context *gin.Context) {
	uid := context.Param("uid")

	err := handler.service.DeleteUser(uid, context)
	if err != nil {
		log.Printf("error with deleting user in firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting user in firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "user deleted successfully in firebase", "id": uid})
}
