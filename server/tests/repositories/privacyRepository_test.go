package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetPrivacy(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock store: %s", err)
	}
	defer db.Close()

	privacys := []types.Privacy{
		{
			ID:      "1",
			Title:   "Title 1",
			Body:    "Body 1",
			Created: time.Now(),
		},
		{
			ID:      "2",
			Title:   "Title 2",
			Body:    "Body 2",
			Created: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "body", "created"})
	for _, privacy := range privacys {
		rows.AddRow(privacy.ID, privacy.Title, privacy.Body, privacy.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Privacy"`).WillReturnRows(rows)

	repo := repository.NewPrivacyRepository(db)
	retrieved, err := repo.GetPrivacy()
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(privacys))
		assert.Equal(test, privacys, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}

func TestCreatePrivacy(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock store: %s", err)
	}
	defer db.Close()

	privacy := &types.Privacy{
		ID:      "1",
		Title:   "Title 1",
		Body:    "Body 1",
		Created: time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Privacy" \(id, title, body, created\) VALUES \(\$1, \$2, \$3, \$4\)`).
		WithArgs(sqlmock.AnyArg(), privacy.Title, privacy.Body, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewPrivacyRepository(db)
	err = repo.CreatePrivacy(privacy)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
