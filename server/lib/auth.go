package lib

import (
	"errors"
	"net/smtp"
)

type Auth interface {
	Start(server *smtp.ServerInfo) (string, []byte, error)
	Next(fromServer []byte, more bool) ([]byte, error)
}

type AuthImpl struct {
	username, password string
}

func NewLoginAuth(username, password string) Auth {
	return &AuthImpl{
		username: username,
		password: password,
	}
}

func (auth *AuthImpl) Start(_ *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (auth *AuthImpl) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(auth.username), nil
		case "Password:":
			return []byte(auth.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}
