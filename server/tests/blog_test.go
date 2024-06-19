package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"

	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("CREATE TABLE Blog").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO Blog").WithArgs(
		sqlmock.AnyArg(), "Test Title", "2023-01-01", "image.jpg", "Test Subheading", "Test Body", "2023-01-01 10:00:00",
	).WillReturnResult(sqlmock.NewResult(1, 1))

	schema := `CREATE TABLE Blog (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        date TEXT,
        mainImage TEXT,
        subheading TEXT,
        body TEXT,
        created DATETIME
    );`

	if _, err := db.Exec(schema); err != nil {
		return nil, nil, fmt.Errorf("failed to create schema: %w", err)
	}

	insertQuery := `INSERT INTO Blog (id, title, date, main_image, subheading, body, created) VALUES (?, ?, ?, ?, ?, ?, ?);`
	if _, err := db.Exec(insertQuery, "Test Title", "2023-01-01", "image.jpg", "Test Subheading", "Test Body", "2023-01-01 10:00:00"); err != nil {
		return nil, nil, fmt.Errorf("failed to insert test data: %w", err)
	}

	return db, mock, nil
}

func initializeHandler(t *testing.T) (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	db, mock, err := setupTestDB(t)
	if err != nil {
		log.Fatalf("Failed to set up test database: %v", err)
	}

	store := stores.NewBlogStore(db)
	s3Client := lib.NewNoOpS3Client()
	service := services.NewBlogService(store, s3Client, "test-folder")
	handler := handlers.NewBlogHandler(service)

	router := gin.Default()
	router.GET("/blogs", handler.GetBlogs)
	router.POST("/blogs", handler.CreateBlog)

	return router, db, mock
}

func TestGetBlogs(t *testing.T) {
	router, db, mock := initializeHandler(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	rows := sqlmock.NewRows([]string{"id", "title", "date", "mainImage", "subheading", "body", "created"}).
		AddRow(1, "Test Title", "2023-01-01", "image.jpg", "Test Subheading", "Test Body", "2023-01-01 10:00:00")
	mock.ExpectQuery("SELECT * FROM Blog").WillReturnRows(rows)

	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(fmt.Sprintf("%s/blogs", server.URL))
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var blogs []types.Blog
	err = json.NewDecoder(resp.Body).Decode(&blogs)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	assert.Len(t, blogs, 1)
	assert.Equal(t, "Test Title", blogs[0].Title)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//func TestCreateBlog(t *testing.T) {
//	router, db := initializeHandler()
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			return
//		}
//	}(db)
//
//	server := httptest.NewServer(router)
//	defer server.Close()
//
//	blog := types.Blog{
//		Title:      "New Blog",
//		Date:       "2024-01-01",
//		MainImage:  "new_image.jpg",
//		Subheading: "New Subheading",
//		Body:       "New Body",
//		Created:    "2024-01-01 12:00:00",
//	}
//	payload, _ := json.Marshal(blog)
//
//	resp, err := http.Post(fmt.Sprintf("%s/blogs", server.URL), "application/json", bytes.NewBuffer(payload))
//	if err != nil {
//		t.Fatalf("Request failed: %v", err)
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			return
//		}
//	}(resp.Body)
//
//	var count int
//	err = db.QueryRow(`SELECT COUNT(*) FROM Blog WHERE Title = ?`, "New Blog").Scan(&count)
//	if err != nil {
//		t.Fatalf("Failed to query database: %v", err)
//	}
//
//	assert.Equal(t, 1, count)
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
