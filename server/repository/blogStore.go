package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"
)

type BlogRepo interface {
	GetBlogs(ctx context.Context) ([]db.Blog, error)
	GetBlogById(ctx context.Context, id string) (*db.Blog, error)
	CreateBlog(ctx context.Context, blog *db.Blog) error
	UpdateBlog(ctx context.Context, id string, blog *db.Blog) error
	DeleteBlog(ctx context.Context, id string) error
}

type BlogRepoImpl struct {
	queries *db.Queries
}

func NewBlogRepo(sql *sql.DB) *BlogRepoImpl {
	queries := db.New(sql)
	return &BlogRepoImpl{
		queries: queries,
	}
}

func (store *BlogRepoImpl) GetBlogs(ctx context.Context) ([]db.Blog, error) {
	blogs, err := store.queries.GetBlogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}

	var result []db.Blog
	for _, blog := range blogs {
		result = append(result, db.Blog{
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

func (store *BlogRepoImpl) GetBlogById(ctx context.Context, id string) (*db.Blog, error) {
	blog, err := store.queries.GetBlogByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}

	return &blog, nil
}

func (store *BlogRepoImpl) CreateBlog(ctx context.Context, blog *db.Blog) error {
	blog.ID = uuid.NewString()
	blog.Created = sql.NullString{String: time.Now().String(), Valid: true}

	params := db.CreateBlogParams{
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

func (store *BlogRepoImpl) UpdateBlog(ctx context.Context, id string, blog *db.Blog) error {
	if blog.MainImage.Valid {
		params := db.UpdateBlogParams{
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
		params := db.UpdateBlogNoImageParams{
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

func (store *BlogRepoImpl) DeleteBlog(ctx context.Context, id string) error {
	err := store.queries.DeleteBlog(ctx, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting blog: %w", err)
	}

	return nil
}
