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

func TestGetExhibition(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	exhibitions := []types.Exhibition{
		{
			ID:       "1",
			Title:    "Exhibit 1",
			Date:     "01/01/24",
			Location: "Dublin",
			Info:     "Stand 1",
			Created:  time.Now(),
		},
		{
			ID:       "2",
			Title:    "Exhibit 2",
			Date:     "01/01/24",
			Location: "Dublin",
			Info:     "Stand 2",
			Created:  time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "date", "location", "info", "created"})
	for _, exhibit := range exhibitions {
		rows.AddRow(exhibit.ID, exhibit.Title, exhibit.Date, exhibit.Location, exhibit.Info, exhibit.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Exhibitions"`).WillReturnRows(rows)

	repo := repository.NewExhibitionRepository(db)
	retrieved, err := repo.GetExhibitions()
	if err != nil {
		test.Fatalf("error occurred when getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(exhibitions))
		assert.Equal(test, exhibitions, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("expectations unfullfilled")
	}
}

func TextCreateExhibition(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	exhibition := &types.Exhibition{
		ID:       "1",
		Title:    "Exhibit 1",
		Date:     "01/01/24",
		Location: "Dublin",
		Info:     "Stand 1",
		Created:  time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Exhibition" \(id, title, data, location, info, created\) 
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
		WithArgs(sqlmock.AnyArg(), exhibition.Title, exhibition.Date, exhibition.Location, exhibition.Info, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewExhibitionRepository(db)
	err = repo.CreateExhibition(exhibition)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
