package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type BlogStore interface {
	GetBlogs() ([]types.Blog, error)
	GetBlogById(id string) (*types.Blog, error)
	CreateBlog(blog *types.Blog) error
	UpdateBlog(id string, blog *types.Blog) error
	DeleteBlog(id string) error
}

type BlogStoreImpl struct {
	database *sql.DB
}

func NewBlogStore(database *sql.DB) *BlogStoreImpl {
	return &BlogStoreImpl{database: database}
}

func (store *BlogStoreImpl) GetBlogs() ([]types.Blog, error) {
	var blogs []types.Blog

	query := `SELECT * FROM "Blog" ORDER BY "created" DESC`
	rows, err := store.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying databse: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var blog types.Blog

		err := rows.Scan(&blog.ID, &blog.Title, &blog.Date, &blog.MainImage, &blog.Subheading, &blog.Body, &blog.Created)
		if err != nil {
			return nil, fmt.Errorf("error occurred while scanning rows: %w", err)
		}

		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
	}

	return blogs, nil
}

func (store *BlogStoreImpl) GetBlogById(id string) (*types.Blog, error) {
	query := `SELECT * FROM "Blog" WHERE "id" = ?`
	row := store.database.QueryRow(query, id)

	var blog types.Blog

	err := row.Scan(&blog.ID, &blog.Title, &blog.Date, &blog.MainImage, &blog.Subheading, &blog.Body, &blog.Created)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error occurred while getting blog: %w", err)
	}

	return &blog, nil
}

func (store *BlogStoreImpl) CreateBlog(blog *types.Blog) error {
	blog.ID = uuid.NewString()
	blog.Created = time.Now().String()

	query := `INSERT INTO "Blog" (id, title, date, main_image, subheading, body, created) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, blog.ID, blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, blog.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating blog: %w", err)
	}

	return nil
}

func (store *BlogStoreImpl) UpdateBlog(id string, blog *types.Blog) error {
	query := `UPDATE "Blog" SET "title" = ?, "date" = ?, "subheading" = ?, "body" = ? WHERE "id" = ?`
	args := []interface{}{blog.Title, blog.Date, blog.Subheading, blog.Body, id}

	if blog.MainImage != "" && blog.MainImage != "null" {
		query = `UPDATE "Blog" SET "title" = ?, "date" = ?, "main_image" = ?, "subheading" = ?, "body" = ? WHERE "id" = ?`
		args = []interface{}{blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, id}
	}

	_, err := store.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error occurred while updating blog: %w", err)
	}

	return nil
}

func (store *BlogStoreImpl) DeleteBlog(id string) error {
	query := `DELETE FROM "Blog" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting blog: %w", err)
	}

	return nil
}
