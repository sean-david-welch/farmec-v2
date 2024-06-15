package services

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"google.golang.org/api/iterator"
)

type AuthService interface {
	Login(context context.Context, idToken string) (string, error)
	Register(context context.Context, userData *types.UserData) error
	GetUsers(context context.Context) ([]*auth.UserRecord, error)
	UpdateUser(uid string, userData *types.UserData, context context.Context) error
	DeleteUser(uid string, context context.Context) error
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

func (service *AuthServiceImpl) GetUsers(context context.Context) ([]*auth.UserRecord, error) {
	authClient, err := service.firebase.App.Auth(context)
	if err != nil {
		return nil, fmt.Errorf("firebase auth initilization failed: %w", err)
	}

	var allUsers []*auth.UserRecord
	pager := iterator.NewPager(authClient.Users(context, ""), 100, "")

	for {
		var users []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			return nil, fmt.Errorf("error listing users: %w", err)
		}
		for _, user := range users {
			allUsers = append(allUsers, &auth.UserRecord{
				UserInfo:      user.UserInfo,
				EmailVerified: user.EmailVerified,
				CustomClaims:  user.CustomClaims,
			})
		}

		if nextPageToken == "" {
			break
		}
	}

	return allUsers, nil
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
	} else {
		err := authClient.SetCustomUserClaims(context, userRecord.UID, map[string]interface{}{"admin": false})
		if err != nil {
			return fmt.Errorf("failed to set custom user claims: %w", err)
		}
	}

	return nil
}

func (service *AuthServiceImpl) UpdateUser(uid string, userData *types.UserData, context context.Context) error {
	authClient, err := service.firebase.App.Auth(context)
	if err != nil {
		return fmt.Errorf("firebase auth initilization failed: %w", err)
	}
	userRecord, err := authClient.UpdateUser(context, uid, (&auth.UserToUpdate{}).Email(userData.Email).Password(userData.Password))
	if err != nil {
		return fmt.Errorf("failed to update user record in firebase: %w", err)
	}
	if userData.Role == "admin" {
		err := authClient.SetCustomUserClaims(context, userRecord.UID, map[string]interface{}{"admin": true})
		if err != nil {
			return fmt.Errorf("failed to set custom user claims: %w", err)
		}
	} else {
		err := authClient.SetCustomUserClaims(context, userRecord.UID, map[string]interface{}{"admin": false})
		if err != nil {
			return fmt.Errorf("failed to set custom user claims: %w", err)
		}
	}
	return nil
}

func (service *AuthServiceImpl) DeleteUser(uid string, context context.Context) error {
	authClient, err := service.firebase.App.Auth(context)
	if err != nil {
		return fmt.Errorf("firebase auth initilization failed: %w", err)
	}
	err = authClient.DeleteUser(context, uid)
	if err != nil {
		return fmt.Errorf("failed to delete user in firebase: %w", err)
	}
	return nil
}
