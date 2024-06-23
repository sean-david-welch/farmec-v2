package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/database"
	"time"
)

type BlogStore interface {
	GetBlogs() ([]database.Blog, error)
	GetBlogById(id string) (*database.Blog, error)
	CreateBlog(blog *database.Blog) error
	UpdateBlog(id string, blog *database.Blog) error
	DeleteBlog(id string) error
}

type BlogStoreImpl struct {
	queries *database.Queries
}

func NewBlogStore(sql *sql.DB) *BlogStoreImpl {
	queries := database.New(sql)
	return &BlogStoreImpl{
		queries: queries,
	}
}

func (store *BlogStoreImpl) GetBlogs() ([]database.Blog, error) {
	ctx := context.Background()
	blogs, err := store.queries.GetBlogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}

	// Convert from the generated database to your database if needed
	var result []database.Blog
	for _, blog := range blogs {
		result = append(result, database.Blog{
			ID:         blog.ID,
			Title:      blog.Title,
			Date:       blog.Date,
			MainImage:  blog.MainImage,
			Subheading: blog.Subheading,
			Body:       blog.Body,
			Created:    blog.Created,
		})
	}

	return result, nil
}

func (store *BlogStoreImpl) GetBlogById(id string) (*database.Blog, error) {
	ctx := context.Background()
	blog, err := store.queries.GetBlogByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}

	return &blog, nil
}

func (store *BlogStoreImpl) CreateBlog(blog *database.Blog) error {
	ctx := context.Background()
	blog.ID = uuid.NewString()
	blog.Created = sql.NullString{String: time.Now().String(), Valid: true}

	params := database.CreateBlogParams{
		ID:         blog.ID,
		Title:      blog.Title,
		Date:       blog.Date,
		MainImage:  blog.MainImage,
		Subheading: blog.Subheading,
		Body:       blog.Body,
		Created:    blog.Created,
	}

	err := store.queries.CreateBlog(ctx, params)
	if err != nil {
		return fmt.Errorf("error occurred while creating blog: %w", err)
	}

	return nil
}

func (store *BlogStoreImpl) UpdateBlog(id string, blog *database.Blog) error {
	ctx := context.Background()

	if blog.MainImage.Valid {
		params := database.UpdateBlogParams{
			Title:      blog.Title,
			MainImage:  blog.MainImage,
			Date:       blog.Date,
			Subheading: blog.Subheading,
			Body:       blog.Body,
			ID:         id,
		}
		err := store.queries.UpdateBlog(ctx, params)
		if err != nil {
			return fmt.Errorf("error occurred while updating blog with image: %w", err)
		}
	} else {
		params := database.UpdateBlogNoImageParams{
			Title:      blog.Title,
			Date:       blog.Date,
			Subheading: blog.Subheading,
			Body:       blog.Body,
			ID:         id,
		}
		err := store.queries.UpdateBlogNoImage(ctx, params)
		if err != nil {
			return fmt.Errorf("error occurred while updating blog without image: %w", err)
		}
	}

	return nil
}

func (store *BlogStoreImpl) DeleteBlog(id string) error {
	ctx := context.Background()

	err := store.queries.DeleteBlog(ctx, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting blog: %w", err)
	}

	return nil
}
