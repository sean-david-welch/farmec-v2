package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetTerms(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock store: %s", err)
	}
	defer db.Close()

	terms := []types.Terms{
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
	for _, term := range terms {
		rows.AddRow(term.ID, term.Title, term.Body, term.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Terms"`).WillReturnRows(rows)

	repo := repository.NewTermsRepository(db)
	retrieved, err := repo.GetTerms()
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(terms))
		assert.Equal(test, terms, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateTerm(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock store: %s", err)
	}
	defer db.Close()

	term := &types.Terms{
		ID:      "1",
		Title:   "Title 1",
		Body:    "Body 1",
		Created: time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Terms" \(id, title, body, created\) VALUES \(\$1, \$2, \$3, \$4\)`).
		WithArgs(sqlmock.AnyArg(), term.Title, term.Body, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewTermsRepository(db)
	err = repo.CreateTerm(term)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
