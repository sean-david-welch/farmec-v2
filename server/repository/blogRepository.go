package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type BlogRepository interface {
	GetBlogs() ([]types.Blog, error)
	GetBlogById(id string) (*types.Blog, error)
	CreateBlog(blog *types.Blog) error
	UpdateBlog(id string, blog *types.Blog) error
	DeleteBlog(id string) error
}

type BlogRepositoryImpl struct {
	database *sql.DB
}

func NewBlogRepository(database *sql.DB) *BlogRepositoryImpl {
	return &BlogRepositoryImpl{database: database}
}

func (repository *BlogRepositoryImpl) GetBlogs() ([]types.Blog, error) {
	var blogs []types.Blog

	query := `SELECT * FROM "Blog" ORDER BY "created" DESC`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying databse: %w", err)
	}
	defer rows.Close()

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

func (repository *BlogRepositoryImpl) GetBlogById(id string) (*types.Blog, error) {
	query := `SELECT * FROM "Blog" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	var blog types.Blog

	err := row.Scan(&blog.ID, &blog.Title, &blog.Date, &blog.MainImage, &blog.Subheading, &blog.Body, &blog.Created)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error occurred while getting blog: %w", err)
	}

	return &blog, nil
}

func (repository *BlogRepositoryImpl) CreateBlog(blog *types.Blog) error {
	blog.ID = uuid.NewString()
	blog.Created = time.Now()

	query := `INSERT INTO "Blog" (id, title, date, main_image, subheading, body, created) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := repository.database.Exec(query, blog.ID, blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, blog.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating blog: %w", err)
	}

	return nil
}

func (repository *BlogRepositoryImpl) UpdateBlog(id string, blog *types.Blog) error {
	query := `UPDATE "Blog" SET "title" = $1, "date" = $2, "subheading" = $3, "body" = $4 WHERE "id" = $5`
	args := []interface{}{blog.Title, blog.Date, blog.Subheading, blog.Body, id}

	if blog.MainImage != "" && blog.MainImage != "null" {
		query = `UPDATE "Blog" SET "title" = $1, "date" = $2, "main_image" = $3, "subheading" = $4, "body" = $5 WHERE "id" = $6`
		args = []interface{}{blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, id}
	}

	_, err := repository.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error occurred while updating blog: %w", err)
	}

	return nil
}

func (repository *BlogRepositoryImpl) DeleteBlog(id string) error {
	query := `DELETE FROM "Blog" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting blog: %w", err)
	}

	return nil
}
