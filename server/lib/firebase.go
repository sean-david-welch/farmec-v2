package lib

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"google.golang.org/api/option"
)

type Firebase struct {
	App *firebase.App
}

func NewFirebase(secrets *config.Secrets) *Firebase {
	opt := option.WithCredentialsJSON([]byte(`{
			"type": "service_account",
			"project_id": "` + secrets.ProjectId + `",
			"private_key_id": "` + secrets.PrivateKeyId + `",
			"private_key": "` + secrets.PrivateKey + `",
			"client_email": "` + secrets.ClientEmail + `",
			"client_id": "` + secrets.ClientId + `",
			"auth_uri": "` + secrets.AuthUri + `",
			"token_uri": "` + secrets.TokenUri + `",
			"auth_provider_x509_cert_url": "` + secrets.AuthProviderX509CertUrl + `",
			"client_x509_cert_url": "` + secrets.ClientX509CertUrl + `"
		}`))

	app, err := firebase.NewApp(context.Background(), nil, opt); if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
	}

	return &Firebase{App: app}
}

func (firebase *Firebase) VerifyToken() gin.HandlerFunc {
	return func(con *gin.Context) {
		cookie, err := con.Cookie("session"); if err != nil {
			con.JSON(http.StatusUnauthorized, gin.H{"error": "no toekn provided"})
			con.Abort()
			return
		}

		client, err := firebase.App.Auth(context.Background()); if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{"error": "error getting auth client"})
			con.Abort()
			return
		}

		decodedToken, err := client.VerifySessionCookie(context.Background(), cookie); if err != nil {
			con.JSON(http.StatusUnauthorized, gin.H{"error": "Error Verifying Token"})
			con.Abort()
			return
		}

		isAdmin := decodedToken.Claims["admin"] == true

		con.Set("isAdmin", isAdmin)
		con.Set("token", decodedToken)

		con.Next()
	}
}