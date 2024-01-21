package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetProduct(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock database: %s", err)
	}
	defer db.Close()

	products := []types.Product{
		{
			ID:           "1",
			MachineID:    "12",
			Name:         "Machine 1",
			ProductImage: "image1.jpg",
			Description:  "machine1Name",
			ProductLink:  "machine1Link",
		},
		{
			ID:           "2",
			MachineID:    "12",
			Name:         "machine 2",
			ProductImage: "image2.jpg",
			Description:  "machine2Name",
			ProductLink:  "machine2Link",
		},
	}

	machineId := products[0].ID

	rows := sqlmock.NewRows([]string{"id", "machineId", "name", "product_image", "description", "product_link"})
	for _, product := range products {
		rows.AddRow(product.ID, product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink)
	}

	mock.ExpectQuery(`SELECT \* FROM "Product"`).WillReturnRows(rows)

	repo := repository.NewProductRepository(db)
	retrieved, err := repo.GetProducts(machineId)
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(products))
		assert.Equal(test, products, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateProduct(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	product := &types.Product{
		ID:           "1",
		MachineID:    "12",
		Name:         "Machine 1",
		ProductImage: "image1.jpg",
		Description:  "machine1Name",
		ProductLink:  "machine1Link",
	}

	mock.ExpectExec(`INSERT INTO "Product" \(id, machineId, name, product_image, description, product_link\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), product.Name, product.ProductImage, product.ProductLink).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewProductRepository(db)
	err = repo.CreateProduct(product)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
