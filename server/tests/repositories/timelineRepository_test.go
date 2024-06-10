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

func TestGetTimelines(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock store: %s", err)
	}
	defer db.Close()

	timelines := []types.Timeline{
		{
			ID:      "1",
			Title:   "Title 1",
			Date:    "01/01/24",
			Body:    "Body 1",
			Created: time.Now(),
		},
		{
			ID:      "2",
			Title:   "Title 2",
			Date:    "01/01/24",
			Body:    "Body 2",
			Created: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "date", "body", "created"})
	for _, timeline := range timelines {
		rows.AddRow(timeline.ID, timeline.Title, timeline.Date, timeline.Body, timeline.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Timeline"`).WillReturnRows(rows)

	repo := repository.NewTimelineRepository(db)
	retrieved, err := repo.GetTimelines()
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(timelines))
		assert.Equal(test, timelines, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateTimeline(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock store: %s", err)
	}
	defer db.Close()

	timeline := &types.Timeline{
		ID:      "1",
		Title:   "Title 1",
		Date:    "01/01/24",
		Body:    "Body 1",
		Created: time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Timeline" \(id, title, date, body, created\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
		WithArgs(sqlmock.AnyArg(), timeline.Title, timeline.Date, timeline.Body, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewTimelineRepository(db)
	err = repo.CreateTimeline(timeline)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
