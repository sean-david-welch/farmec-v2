package lib

import (
	"errors"
	"net/smtp"
)

type EmailAuth interface {
	Start(_ *smtp.ServerInfo) (string, []byte, error)
	Next(fromServer []byte, more bool) ([]byte, error)
}

type EmailAuthImpl struct {
	username, password string
}

func NewLoginAuth(username, password string) EmailAuth {
	return &EmailAuthImpl{
		username: username,
		password: password,
	}
}

func (auth *EmailAuthImpl) Start(_ *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (auth *EmailAuthImpl) Next(fromServer []byte, more bool) ([]byte, error) {
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
