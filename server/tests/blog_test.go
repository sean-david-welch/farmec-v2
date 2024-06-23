package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
)

type BlogTestSuite struct {
	suite.Suite
	db     *sql.DB
	mock   sqlmock.Sqlmock
	router *gin.Engine
}

func (suite *BlogTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	require.NoError(suite.T(), err)

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
	require.NoError(suite.T(), err)

	suite.db = db
	suite.mock = mock

	store := stores.NewBlogStore(db)
	s3Client := lib.NewNoOpS3Client()
	service := services.NewBlogService(store, s3Client, "test-folder")
	handler := handlers.NewBlogHandler(service)

	suite.router = gin.Default()
	suite.router.GET("/blogs", handler.GetBlogs)
	suite.router.POST("/blogs", handler.CreateBlog)
}

func (suite *BlogTestSuite) TearDownTest() {
	err := suite.db.Close()
	require.NoError(suite.T(), err)
}

func (suite *BlogTestSuite) TestGetBlogs() {
	rows := sqlmock.NewRows([]string{"id", "title", "date", "main_image", "subheading", "body", "created"}).
		AddRow("1", "Test Title", "2023-01-01", "image.jpg", "Test Subheading", "Test Body", "2023-01-01 10:00:00")
	suite.mock.ExpectQuery("SELECT id, title, date, main_image, subheading, body, created FROM Blog").WillReturnRows(rows)

	server := httptest.NewServer(suite.router)
	defer server.Close()

	resp, err := http.Get(fmt.Sprintf("%s/blogs", server.URL))
	require.NoError(suite.T(), err)
	defer func() {
		err := resp.Body.Close()
		require.NoError(suite.T(), err)
	}()

	require.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	var blogs []types.Blog
	err = json.NewDecoder(resp.Body).Decode(&blogs)
	require.NoError(suite.T(), err)

	require.Len(suite.T(), blogs, 1)
	require.Equal(suite.T(), "Test Title", blogs[0].Title)

	err = suite.mock.ExpectationsWereMet()
	require.NoError(suite.T(), err)
}

func TestBlogTestSuite(t *testing.T) {
	suite.Run(t, new(BlogTestSuite))
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
