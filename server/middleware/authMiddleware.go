package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
)

type AuthMiddleware struct {
	FirebaseService *lib.Firebase
}

func NewAuthMiddleware(firebaseService *lib.Firebase) *AuthMiddleware {
	return &AuthMiddleware{
		FirebaseService: firebaseService,
	}
}

func (middleware *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := context.Cookie("access_token")
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, No token provided"})
			return
		}

		decodedToken, isAdmin, err := middleware.FirebaseService.VerifyToken(token)
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, Invalid Token"})
			return
		}

		context.Set("decodedToken", decodedToken)
		context.Set("isAdmin", isAdmin)

		context.Next()
	}
}
