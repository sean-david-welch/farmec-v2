package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/services"
)

type BlogTestSuite struct {
	suite.Suite
	db     *sql.DB
	mock   sqlmock.Sqlmock
	router *gin.Engine
}

func (suite *BlogTestSuite) SetupTest() {
	database, mock, err := sqlmock.New()
	require.NoError(suite.T(), err)

	mock.ExpectExec(`CREATE TABLE Blog \(
        id TEXT PRIMARY KEY,
        title TEXT,
        date TEXT,
        main_image TEXT,
        subheading TEXT,
        body TEXT,
        created TEXT
    \);`).WillReturnResult(sqlmock.NewResult(0, 0))

	_, err = database.Exec(`CREATE TABLE Blog (
        id TEXT PRIMARY KEY,
        title TEXT,
        date TEXT,
        main_image TEXT,
        subheading TEXT,
        body TEXT,
        created TEXT
    );`)
	require.NoError(suite.T(), err)

	suite.db = database
	suite.mock = mock

	store := repository.NewBlogStore(database)
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

	var blogs []db.Blog
	err = json.NewDecoder(resp.Body).Decode(&blogs)
	require.NoError(suite.T(), err)

	require.Len(suite.T(), blogs, 1)
	require.Equal(suite.T(), "Test Title", blogs[0].Title)
	require.Equal(suite.T(), "2023-01-01", blogs[0].Date.String)

	err = suite.mock.ExpectationsWereMet()
	require.NoError(suite.T(), err)
}

func (suite *BlogTestSuite) TestCreateBlog() {
	blog := db.Blog{
		ID:         "",
		Title:      "New Blog",
		Date:       sql.NullString{String: "2024-06-23", Valid: true},
		MainImage:  sql.NullString{String: "image.jpg", Valid: true},
		Subheading: sql.NullString{String: "This is a subheading", Valid: true},
		Body:       sql.NullString{String: "This is the body of the blog.", Valid: true},
		Created:    sql.NullString{String: "2024-06-23 10:00:00", Valid: true},
	}

	payload, err := json.Marshal(blog)
	require.NoError(suite.T(), err)

	suite.mock.ExpectExec("INSERT INTO Blog").
		WithArgs(sqlmock.AnyArg(), blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, blog.Created).
		WillReturnResult(sqlmock.NewResult(1, 1))

	server := httptest.NewServer(suite.router)
	defer server.Close()

	resp, err := http.Post(fmt.Sprintf("%s/blogs", server.URL), "application/json", bytes.NewBuffer(payload))
	require.NoError(suite.T(), err)
	defer func() {
		err := resp.Body.Close()
		require.NoError(suite.T(), err)
	}()

	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode)

	err = suite.mock.ExpectationsWereMet()
	require.NoError(suite.T(), err)

	var count int
	err = suite.db.QueryRow(`SELECT COUNT(*) FROM Blog WHERE title = ?`, "New Blog").Scan(&count)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, count)
}

func TestBlogTestSuite(t *testing.T) {
	suite.Run(t, new(BlogTestSuite))
}
