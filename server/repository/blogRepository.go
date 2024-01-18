package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type BlogRepository struct {
	database *sql.DB
}

func NewBlogRepository(database *sql.DB) *BlogRepository {
	return &BlogRepository{database: database}
}

func(repository *BlogRepository) GetBlogs() ([]types.Blog, error) {
	var blogs []types.Blog

	query := `SELECT * FROM "Blog"`
	rows, err := repository.database.Query(query); if err != nil {
		return nil, fmt.Errorf("error occurred while querying databse: %w", err)
	}
	defer rows.Close()

	for rows.Next(){
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

func(repository *BlogRepository) GetBlogById(id string) (*types.Blog, error) {
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

func(repository *BlogRepository) CreateBlog(blog *types.Blog) error {
	blog.ID = uuid.NewString()
	blog.Created = time.Now()

	query := `INSERT INTO "Blog" (id, title, date, mainImage, subHeading, body, created) VALUES ($1, $2, $3, $4, $5, $6, $7)`

    _, err := repository.database.Exec(query, blog.ID, blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, blog.Created)
	if err != nil {
		return fmt.Errorf("error occurred while creating blog: %w", err)
	}

	return nil
}

func(repository *BlogRepository) UpdateBlog(id string, blog *types.Blog) error {
	query := `UPDATE "Blog" SET "title" = $1, "date" = $2, "mainImage" = $3, "subHeading" = $4, "body" = $5 WHERE "id" = $6`

    _, err := repository.database.Exec(query, blog.Title, blog.Date, blog.MainImage, blog.Subheading, blog.Body, blog.Created, id)
	if err != nil {
		return fmt.Errorf("error occurred while updating blog: %w", err)
	}
	
	return nil
}

func(repository *BlogRepository) DeleteBlog(id string) error {
	query := `DELETE FROM "Blog" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id); if err != nil {
		return fmt.Errorf("error occurred while deleting blog: %w", err)
	}

	return nil
}

