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
	log.Printf("Incoming login request from IP: %s, Path: %s", context.ClientIP(), context.Request.URL.Path)

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

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("session", sessionCookie, 72*3600, "/", "", true, true)

	log.Printf("Cookie set for session: %s", sessionCookie)
	log.Printf("Setting cookie: Name=%s; Value=%s; MaxAge=%d; Path=%s; Domain=%s; Secure=%t; HttpOnly=%t; SameSite=None",
		"session", sessionCookie, 72*3600, "/", "", true, true)

	context.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (controller *AuthController) GetUsers(context *gin.Context) {
	users, err := controller.service.GetUsers(context)
	if err != nil {
		log.Printf("error fetching users from firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching users from firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
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

func (controller *AuthController) UpdateUser(context *gin.Context) {
	uid := context.Param("uid")
	var userData types.UserData

	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while updating user in firebase", "details": err.Error()})
		return
	}

	if err := controller.service.UpdateUser(uid, &userData, context); err != nil {
		log.Printf("error while updating user in firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating user in firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user successfully updated in firebase"})
}

func (controller *AuthController) DeleteUser(context *gin.Context) {
	uid := context.Param("uid")

	err := controller.service.DeleteUser(uid, context)
	if err != nil {
		log.Printf("error with deleting user in firebase: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting user in firebase", "details": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "user deleted successfully in firebase", "id": uid})
}
