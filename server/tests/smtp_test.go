package tests

import (
	"crypto/tls"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/smtp"
	"testing"

	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/stretchr/testify/assert"
)

func TestSMTPAuth(t *testing.T) {
	secrets, err := lib.NewSecrets()
	if err != nil {
		return
	}
	emailUser := secrets.EmailUser
	emailPass := secrets.EmailPass

	if emailUser == "" || emailPass == "" {
		t.Skip("Skipping test: EMAIL_USER or EMAIL_PASS environment variables not set")
	}

	tests := []struct {
		name    string
		user    string
		pass    string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid Credentials",
			user:    emailUser,
			pass:    emailPass,
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Invalid Password",
			user:    emailUser,
			pass:    "wrongpassword",
			wantErr: true,
			errMsg:  "auth failed",
		},
		{
			name:    "Empty Credentials",
			user:    "",
			pass:    "",
			wantErr: true,
			errMsg:  "auth failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := lib.NewLoginAuth(tt.user, tt.pass)

			// Connect to the SMTP Server
			c, err := smtp.Dial("smtp.office365.com:587")
			if err != nil {
				t.Fatalf("dial failed: %v", err)
			}
			defer func(c *smtp.Client) {
				err := c.Close()
				if err != nil {
					return
				}
			}(c)

			// Send EHLO
			if err = c.Hello("localhost"); err != nil {
				t.Fatalf("hello failed: %v", err)
			}

			// Start TLS
			if err = c.StartTLS(&tls.Config{ServerName: "smtp.office365.com"}); err != nil {
				t.Fatalf("starttls failed: %v", err)
			}

			// Auth
			err = c.Auth(auth)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Test the full SMTP client implementation
func TestSMTPClientImpl(t *testing.T) {
	// Get credentials from environment variables for testing
	secrets, err := lib.NewSecrets()
	if err != nil {
		return
	}
	emailUser := secrets.EmailUser
	emailPass := secrets.EmailPass

	if emailUser == "" || emailPass == "" {
		t.Skip("Skipping test: EMAIL_USER or EMAIL_PASS environment variables not set")
	}

	client := lib.NewSTMPClient(secrets)

	// Test SetupSMTPClient
	t.Run("SetupSMTPClient", func(t *testing.T) {
		smtpClient, err := client.SetupSMTPClient()
		assert.NoError(t, err)
		assert.NotNil(t, smtpClient)

		if smtpClient != nil {
			err := smtpClient.Quit()
			if err != nil {
				return
			}
		}
	})

	// Test SendFormNotification
	t.Run("SendFormNotification", func(t *testing.T) {
		testData := &types.EmailData{
			Name:    "Test User",
			Email:   "test@example.com",
			Message: "This is a test message",
		}

		err := client.SendFormNotification(testData, "Test")
		assert.NoError(t, err)
	})
}

// Test error cases
func TestSMTPClientImplErrors(t *testing.T) {
	// Test with invalid credentials
	secrets := &lib.Secrets{
		EmailUser: "invalid@example.com",
		EmailPass: "wrongpassword",
	}

	client := lib.NewSTMPClient(secrets)

	t.Run("SetupSMTPClient with invalid credentials", func(t *testing.T) {
		smtpClient, err := client.SetupSMTPClient()
		assert.Error(t, err)
		assert.Nil(t, smtpClient)
	})

	t.Run("SendFormNotification with invalid credentials", func(t *testing.T) {
		testData := &types.EmailData{
			Name:    "Test User",
			Email:   "test@example.com",
			Message: "This is a test message",
		}

		err := client.SendFormNotification(testData, "Test")
		assert.Error(t, err)
	})
}
