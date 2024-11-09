package lib

import (
	"net/smtp"
)

type LoginAuth struct {
	username string
	password string
	host     string
}

func (a *LoginAuth) Start(_ *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		return []byte(a.password), nil
	}
	return nil, nil
}

func NewLoginAuth(username, password string) smtp.Auth {
	return &LoginAuth{
		username: username,
		password: password,
		host:     "smtp.office365.com",
	}
}
