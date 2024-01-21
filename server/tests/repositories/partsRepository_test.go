package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetParts(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database")
	}
	defer db.Close()

	parts := []types.Sparepart{
		{
			ID:             "1",
			SupplierID:     "12",
			Name:           "Part 1",
			PartsImage:     "image1.jpg",
			SparePartsLink: "www.google.com",
		},
		{
			ID:             "2",
			SupplierID:     "12",
			Name:           "Part 2",
			PartsImage:     "image2.jpg",
			SparePartsLink: "www.google.com",
		},
	}

	supplierId := parts[0].SupplierID

	rows := sqlmock.NewRows([]string{"id", "supplierId", "name", "parts_image", "spare_parts_links"})
	for _, part := range parts {
		rows.AddRow(part.ID, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink)
	}

	mock.ExpectQuery(`SELECT \* FROM "SpareParts"`).WillReturnRows(rows)

	repo := repository.NewPartsRepository(db)
	retrieved, err := repo.GetParts(supplierId)
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(parts))
		assert.Equal(test, parts, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}

func TestCreatePart(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database")
	}
	defer db.Close()

	part := &types.Sparepart{
		ID:             "1",
		SupplierID:     "12",
		Name:           "Part 1",
		PartsImage:     "image1.jpg",
		SparePartsLink: "www.google.com",
	}

	mock.ExpectExec(`INSERT INTO "SpareParts" \(id, supplierId, name, parts_image, spare_parts_link\)
		VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
		WithArgs(sqlmock.AnyArg(), part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewPartsRepository(db)
	err = repo.CreatePart(part)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
