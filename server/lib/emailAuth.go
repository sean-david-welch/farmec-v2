package lib

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

type Office365Auth struct {
	username string
	password string
}

func (a *Office365Auth) Start(_ *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *Office365Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(base64.StdEncoding.EncodeToString([]byte(a.username))), nil
		case "Password:":
			return []byte(base64.StdEncoding.EncodeToString([]byte(a.password))), nil
		default:
			return nil, fmt.Errorf("unexpected server challenge: %s", string(fromServer))
		}
	}
	return nil, nil
}

func NewOffice365Auth(username, password string) smtp.Auth {
	return &Office365Auth{
		username: username,
		password: password,
	}
}
