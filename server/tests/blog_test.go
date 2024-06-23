package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
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

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	schema := `CREATE TABLE Blog (
		id TEXT PRIMARY KEY,
		title TEXT,
		date TEXT,
		main_image TEXT,
		subheading TEXT,
		body TEXT,
		created TEXT
	);`

	_, err = db.Exec(schema)
	require.NoError(t, err)

	return db, mock
}

func initializeHandler(t *testing.T) (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	db, mock := setupTestDB(t)

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
	defer db.Close()

	// Mock the expected query and result
	rows := sqlmock.NewRows([]string{"id", "title", "date", "main_image", "subheading", "body", "created"}).
		AddRow("1", "Test Title", "2023-01-01", "image.jpg", "Test Subheading", "Test Body", "2023-01-01 10:00:00")
	mock.ExpectQuery("SELECT id, title, date, main_image, subheading, body, created FROM Blog").WillReturnRows(rows)

	// Create a test server
	server := httptest.NewServer(router)
	defer server.Close()

	// Perform a GET request to /blogs
	resp, err := http.Get(fmt.Sprintf("%s/blogs", server.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse the response body
	var blogs []types.Blog
	err = json.NewDecoder(resp.Body).Decode(&blogs)
	require.NoError(t, err)

	// Validate the response
	require.Len(t, blogs, 1)
	assert.Equal(t, "Test Title", blogs[0].Title)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
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
