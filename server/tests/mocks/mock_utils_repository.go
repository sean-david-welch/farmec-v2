package mocks

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func InitMockDatabase(test *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		test.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}
