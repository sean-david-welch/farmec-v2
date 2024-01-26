package lib

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/sean-david-welch/farmec-v2/server/config"
	"google.golang.org/api/option"
)

type Firebase struct {
	App *firebase.App
}

func NewFirebase(secrets *config.Secrets) (*Firebase, error) {
	privateKey := strings.ReplaceAll(secrets.PrivateKey, "\n", "\\n")

	credentialsJSON := `{
			"type": "service_account",
			"project_id": "` + secrets.ProjectId + `",
			"private_key_id": "` + secrets.PrivateKeyId + `",
			"private_key": "` + privateKey + `",
			"client_email": "` + secrets.ClientEmail + `",
			"client_id": "` + secrets.ClientId + `",
			"auth_uri": "` + secrets.AuthUri + `",
			"token_uri": "` + secrets.TokenUri + `",
			"auth_provider_x509_cert_url": "` + secrets.AuthProviderX509CertUrl + `",
			"client_x509_cert_url": "` + secrets.ClientX509CertUrl + `"
	}`

	log.Printf("firebase credentials: %v", credentialsJSON)
	opt := option.WithCredentialsJSON([]byte(credentialsJSON))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Failed to initialize Firebase: %v", err)
		return nil, err
	}

	return &Firebase{App: app}, nil
}

func (firebase *Firebase) VerifyToken(token string) (*auth.Token, bool, error) {
	context := context.Background()

	authClient, err := firebase.App.Auth(context)
	if err != nil {
		return nil, false, err
	}

	decodedToken, err := authClient.VerifyIDToken(context, token)
	if err != nil {
		return nil, false, err
	}

	isAdmin := false
	if adminClaim, ok := decodedToken.Claims["admin"]; ok {
		isAdmin, _ = adminClaim.(bool)
	}

	return decodedToken, isAdmin, nil
}
