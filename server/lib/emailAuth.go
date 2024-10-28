package lib

import (
	"net/smtp"
)

type EmailAuth interface {
	Start(_ *smtp.ServerInfo) (string, []byte, error)
	Next(fromServer []byte, more bool) ([]byte, error)
}

type EmailAuthImpl interface {
	smtp.Auth
}

func NewLoginAuth(username, password string) EmailAuth {
	return smtp.PlainAuth(
		"",
		username,
		password,
		"smtp.office365.com",
	)
}
