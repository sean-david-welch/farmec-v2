package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetLineItems(t *testing.T) {
    db, mock, err := sqlmock.New(); if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    lineItems := []types.LineItem{
        {ID: "1", Name: "Item 1", Price: 10.99, Image: "image1.jpg"},
        {ID: "2", Name: "Item 2", Price: 20.99, Image: "image2.jpg"},
    }

    rows := sqlmock.NewRows([]string{"id", "name", "price", "image"})
    for _, item := range lineItems {
        rows.AddRow(item.ID, item.Name, item.Price, item.Image)
    }

	mock.ExpectQuery(`SELECT \* FROM "LineItems"`).WillReturnRows(rows)

    repo := repository.NewLineItemRepository(db)
    retrievedItems, err := repo.GetLineItems()

    assert.NoError(t, err)
    if err == nil {
        assert.Len(t, retrievedItems, len(lineItems))
        for i, item := range retrievedItems {
			assert.Equal(t, lineItems[i].ID, item.ID)
			assert.Equal(t, lineItems[i].Name, item.Name)
			assert.Equal(t, lineItems[i].Price, item.Price)
			assert.Equal(t, lineItems[i].Image, item.Image)
        }
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCreateLineItem(t *testing.T) {
    db, mock, err := sqlmock.New(); if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    lineItem := &types.LineItem{Name: "Item 1", Price: 10.99, Image: "image1.jpg"}

    mock.ExpectExec(`INSERT INTO "Lineitems" \(id, name, price, image\) VALUES \(\$1, \$2, \$3, \$4\)`).
        WithArgs(sqlmock.AnyArg(), lineItem.Name, lineItem.Price, lineItem.Image).
        WillReturnResult(sqlmock.NewResult(1, 1))

    repo := repository.NewLineItemRepository(db)
    err = repo.CreateLineItem(lineItem)

    assert.NoError(t, err)

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}
