package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/lib"
)

type AdminMiddleware struct {
	FirebaseService *lib.Firebase
}

func NewAdminMiddleware(firebaseService *lib.Firebase) *AdminMiddleware {
	return &AdminMiddleware{
		FirebaseService: firebaseService,
	}
}

func (middleware *AdminMiddleware) Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("AdminMiddleware triggered")

		token, err := context.Cookie("session"); if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, No token provided"})
            return
		}

		decodedToken, isAdmin, err := middleware.FirebaseService.VerifyToken(token); if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized, Invalid Token"})
			return
		}

        if !isAdmin {
            context.AbortWithStatusJSON(403, gin.H{"error": "Forbidden, requires admin privileges"})
            return
        }

		context.Set("decodedToken", decodedToken)

		context.Next()
	}
}