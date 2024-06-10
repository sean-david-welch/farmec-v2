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

func TestGetBlogs(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to initialize mock store: %s", err)
	}
	defer db.Close()

	blogs := []types.Blog{
		{ID: "1", Title: "Blog 1", Date: "17/01/24", MainImage: "image.jpg", Subheading: "Subheading 1", Body: "Body 1", Created: time.Now()},
		{ID: "2", Title: "Blog 2", Date: "17/01/24", MainImage: "image.jpg", Subheading: "Subheading 2", Body: "Body 2", Created: time.Now()},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "date", "main_image", "subheading", "body", "created"})
	for _, blog := range blogs {
		rows.AddRow(blog.ID, blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, blog.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Blog"`).WillReturnRows(rows)

	repo := repository.NewBlogRepository(db)
	retrievedBlogs, err := repo.GetBlogs()
	if err != nil {
		test.Fatalf("error occurred when getting blogs")
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrievedBlogs, len(blogs))
		assert.Equal(test, blogs, retrievedBlogs)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("expectations unfullfilled: %s", err)
	}
}

func TestCreateBlog(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock store: %s", err)
	}
	defer db.Close()

	blog := &types.Blog{Title: "Blog 1", Date: "17/01/24", MainImage: "image.jpg", Subheading: "Subheading 1", Body: "Body 1", Created: time.Now()}
	mock.ExpectExec(`INSERT INTO "Blog" \(id, title, date, main_image, subheading, body, created\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)`).
		WithArgs(sqlmock.AnyArg(), blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewBlogRepository(db)
	err = repo.CreateBlog(blog)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}
