package services

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type AuthService interface {
	Login(context context.Context, idToken string) (string, error)
	Register(context context.Context, userData *types.UserData) error
}

type AuthServiceImpl struct {
	firebase *lib.Firebase
}

func NewAuthService(firebase *lib.Firebase) *AuthServiceImpl {
	return &AuthServiceImpl{firebase: firebase}
}

func (service *AuthServiceImpl) Login(context context.Context, idToken string) (string, error) {

	authClient, err := service.firebase.App.Auth(context)
	if err != nil {
		return "", fmt.Errorf("firebase auth initilization failed: %w", err)
	}

	_, err = authClient.VerifyIDToken(context, idToken)
	if err != nil {
		return "", fmt.Errorf("invalid firebase token: %w", err)

	}

	expiresIn := time.Hour * 72
	sessionCookie, err := authClient.SessionCookie(context, idToken, expiresIn)
	if err != nil {
		return "", fmt.Errorf("failed to create session cookie: %w", err)
	}

	return sessionCookie, nil
}

func (service *AuthServiceImpl) Register(context context.Context, userData *types.UserData) error {
	authClient, err := service.firebase.App.Auth(context)
	if err != nil {
		return fmt.Errorf("firebase auth initilization failed: %w", err)
	}

	userRecord, err := authClient.CreateUser(context, (&auth.UserToCreate{}).
		Email(userData.Email).Password(userData.Password))
	if err != nil {
		return fmt.Errorf("failed to create user in firebase: %w", err)
	}

	if userData.Role == "admin" {
		err := authClient.SetCustomUserClaims(context, userRecord.UID, map[string]interface{}{"admin": true})
		if err != nil {
			return fmt.Errorf("failed to set custom user claims: %w", err)
		}
	}

	return nil
}
