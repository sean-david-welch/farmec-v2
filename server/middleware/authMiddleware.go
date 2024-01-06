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
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("session"); if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, No token provided"})
            return
		}

		decodedToken, isAdmin, err := middleware.FirebaseService.VerifyToken(token); if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, Invalid Token"})
			return
		}

		ctx.Set("decodedToken", decodedToken)
		ctx.Set("isAdmin", isAdmin)

		ctx.Next()
	}
}