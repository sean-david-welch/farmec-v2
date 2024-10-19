package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type BlogService interface {
	GetBlogs(ctx context.Context) ([]types.Blog, error)
	GetBlogsByID(ctx context.Context, id string) (*types.Blog, error)
	CreateBlog(ctx context.Context, blog *db.Blog) (*types.ModelResult, error)
	UpdateBlog(ctx context.Context, id string, blog *db.Blog) (*types.ModelResult, error)
	DeleteBlog(ctx context.Context, id string) error
}

type BlogServiceImpl struct {
	store    repository.BlogRepo
	s3Client lib.S3Client
	folder   string
}

func NewBlogService(store repository.BlogRepo, s3Client lib.S3Client, folder string) *BlogServiceImpl {
	return &BlogServiceImpl{store: store, s3Client: s3Client, folder: folder}
}

func (service *BlogServiceImpl) GetBlogs(ctx context.Context) ([]types.Blog, error) {
	blogs, err := service.store.GetBlogs(ctx)
	if err != nil {
		return nil, err
	}

	var result []types.Blog
	for _, blog := range blogs {
		result = append(result, lib.SerializeBlog(blog))
	}
	return result, nil
}

func (service *BlogServiceImpl) GetBlogsByID(ctx context.Context, id string) (*types.Blog, error) {
	blog, err := service.store.GetBlogById(ctx, id)
	if err != nil {
		return nil, err
	}
	result := lib.SerializeBlog(*blog)

	return &result, nil
}

func (service *BlogServiceImpl) CreateBlog(ctx context.Context, blog *db.Blog) (*types.ModelResult, error) {
	image := blog.MainImage

	if !image.Valid {
		return nil, errors.New("image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	blog.MainImage = sql.NullString{String: imageUrl, Valid: true}

	if err := service.store.CreateBlog(ctx, blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *BlogServiceImpl) UpdateBlog(ctx context.Context, id string, blog *db.Blog) (*types.ModelResult, error) {
	image := blog.MainImage

	var presignedUrl, imageUrl string
	var err error

	if image.Valid {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		blog.MainImage = sql.NullString{String: imageUrl, Valid: true}
	}

	if err := service.store.UpdateBlog(ctx, id, blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     image.String,
	}

	return result, nil
}

func (service *BlogServiceImpl) DeleteBlog(ctx context.Context, id string) error {
	blog, err := service.store.GetBlogById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(blog.MainImage.String); err != nil {
		return err
	}

	if err := service.store.DeleteBlog(ctx, id); err != nil {
		return err
	}

	return nil
}
