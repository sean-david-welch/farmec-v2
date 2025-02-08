package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
)

type AdminAuthenticator interface {
	RouteMiddleware() gin.HandlerFunc
	ViewMiddleware() gin.HandlerFunc
	GetIsAdmin(context *gin.Context) bool
	GetIsAuthenticated(context *gin.Context) bool
}

type AdminMiddleware struct {
	firebase *lib.Firebase
}

func NewAdminMiddleware(firebase *lib.Firebase) *AdminMiddleware {
	return &AdminMiddleware{
		firebase: firebase,
	}
}

func (middleware *AdminMiddleware) RouteMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Cookie("access_token")
		if err != nil {
			context.Set("isAuthenticated", false)
			context.Set("isAdmin", false)
			log.Printf("error occurred no cookie provided: %s", err)
			context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, No cookie provided"})
			return
		}

		decodedToken, isAdmin, err := middleware.firebase.VerifyToken(cookie)
		if err != nil {
			context.Set("isAuthenticated", false)
			context.Set("isAdmin", false)
			log.Printf("error occurred when verifying cookie: %s", err)
			context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, Invalid Token"})
			return
		}

		context.Set("isAuthenticated", true)
		context.Set("isAdmin", isAdmin)

		if !isAdmin {
			log.Printf("error user is not admin: %s", err)
			context.AbortWithStatusJSON(403, gin.H{"error": "Forbidden, requires admin privileges"})
			return
		}

		context.Set("decodedToken", decodedToken)
		context.Next()
	}
}

func (middleware *AdminMiddleware) ViewMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Cookie("access_token")
		if err != nil {
			context.Set("isAuthenticated", false)
			context.Set("isAdmin", false)
			context.Next()
			return
		}

		_, isAdmin, err := middleware.firebase.VerifyToken(cookie)
		context.Set("isAuthenticated", err == nil)
		context.Set("isAdmin", err == nil && isAdmin)
		context.Next()
	}
}

func (middleware *AdminMiddleware) GetIsAdmin(context *gin.Context) bool {
	isAdmin, exists := context.Get("isAdmin")
	if !exists {
		return false
	}
	return isAdmin.(bool)
}

func (middleware *AdminMiddleware) GetIsAuthenticated(context *gin.Context) bool {
	isAuthenticated, exists := context.Get("isAuthenticated")
	if !exists {
		return false
	}
	return isAuthenticated.(bool)
}
