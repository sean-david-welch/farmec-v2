package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
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

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	schema := `CREATE TABLE Blog (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Title TEXT,
		Date TEXT,
		MainImage TEXT,
		Subheading TEXT,
		Body TEXT,
		Created DATETIME
	);
	`
	_, err = db.Exec(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	_, err = db.Exec(`
	INSERT INTO Blog (Title, Date, MainImage, Subheading, Body, Created)
	VALUES ('Test Title', '2023-01-01', 'image.jpg', 'Test Subheading', 'Test Body', '2023-01-01 10:00:00');
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to insert test data: %w", err)
	}

	return db, nil
}

func initializeHandler() (*gin.Engine, *sql.DB) {
	db, err := setupTestDB()
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	store := stores.NewBlogStore(db)
	s3Client := lib.NewNoOpS3Client()
	service := services.NewBlogService(store, s3Client, "test-folder")
	handler := handlers.NewBlogHandler(service)

	router := gin.Default()
	router.GET("/blogs", handler.GetBlogs)
	router.POST("/blogs", handler.CreateBlog)

	return router, db
}

func TestGetBlogs(t *testing.T) {
	router, db := initializeHandler()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

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
}

func TestCreateBlog(t *testing.T) {
	router, db := initializeHandler()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	server := httptest.NewServer(router)
	defer server.Close()

	blog := types.Blog{
		Title:      "New Blog",
		Date:       "2024-01-01",
		MainImage:  "new_image.jpg",
		Subheading: "New Subheading",
		Body:       "New Body",
		Created:    "2024-01-01 12:00:00",
	}
	payload, _ := json.Marshal(blog)

	resp, err := http.Post(fmt.Sprintf("%s/blogs", server.URL), "application/json", bytes.NewBuffer(payload))
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

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM Blog WHERE Title = ?`, "New Blog").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query database: %v", err)
	}

	assert.Equal(t, 1, count)
}
