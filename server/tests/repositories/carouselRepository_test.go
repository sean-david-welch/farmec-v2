package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCarousels(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test); if err != nil {
		test.Fatalf("failed to initialise mock database")
	}
	defer db.Close()

	carousels := []types.Carousel{
		{ID: "1", Name: "Carousel 1", Image: "image1.jpg"},
		{ID: "2", Name: "Carousel 2", Image: "image2.jpg"},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "image"})
	for _, carousel := range carousels {
		rows.AddRow(carousel.ID, carousel.Name, carousel.Image)
	}

	mock.ExpectQuery(`SELECT \* FROM "Carousel"`).WillReturnRows(rows)

	repo := repository.NewCarouselRepository(db)
	retrievedCarousels, err := repo.GetCarousels(); if err != nil {
		test.Errorf("error when getting carousels: %s", err)
	}

	assert.NoError(test, err); if err == nil {
		assert.Len(test, retrievedCarousels, len(carousels))
		assert.Equal(test, carousels, retrievedCarousels)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateCarousel(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test); if err != nil {
		test.Fatalf("failed to init mock database")
	}
	defer db.Close()

	carousel := &types.Carousel{Name: "Carousel 1", Image: "image1.jpg"}	
	mock.ExpectExec(`INSERT INTO "Carousel" \(id, name, image\) VALUES \(\$1, \$2, \$3\)`).
		WithArgs(sqlmock.AnyArg(), carousel.Name, carousel.Image).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewCarouselRepository(db)	
	err = repo.CreateCarousel(carousel)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}