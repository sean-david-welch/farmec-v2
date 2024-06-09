package utils

import (
	"errors"
	"net/smtp"
)

type LoginAuth interface {
	Start(server *smtp.ServerInfo) (string, []byte, error)
	Next(fromServer []byte, more bool) ([]byte, error)
}

type LoginAuthImpl struct {
	username, password string
}

func NewLoginAuth(username, password string) LoginAuth {
	return &LoginAuthImpl{
		username: username,
		password: password,
	}
}

func (auth *LoginAuthImpl) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (auth *LoginAuthImpl) Next(fromServer []byte, more bool) ([]byte, error) {
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
