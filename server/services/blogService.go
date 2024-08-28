package services

import (
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type BlogService interface {
	GetBlogs() ([]db.Blog, error)
	GetBlogsByID(id string) (*db.Blog, error)
	CreateBlog(blog *db.Blog) (*types.ModelResult, error)
	UpdateBlog(id string, blog *db.Blog) (*types.ModelResult, error)
	DeleteBlog(id string) error
}

type BlogServiceImpl struct {
	store    store.BlogStore
	s3Client lib.S3Client
	folder   string
}

func NewBlogService(store store.BlogStore, s3Client lib.S3Client, folder string) *BlogServiceImpl {
	return &BlogServiceImpl{store: store, s3Client: s3Client, folder: folder}
}

func (service *BlogServiceImpl) GetBlogs() ([]db.Blog, error) {
	blogs, err := service.store.GetBlogs()
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (service *BlogServiceImpl) GetBlogsByID(id string) (*db.Blog, error) {
	blog, err := service.store.GetBlogById(id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (service *BlogServiceImpl) CreateBlog(blog *db.Blog) (*types.ModelResult, error) {
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

	if err := service.store.CreateBlog(blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *BlogServiceImpl) UpdateBlog(id string, blog *db.Blog) (*types.ModelResult, error) {
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

	if err := service.store.UpdateBlog(id, blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     image.String,
	}

	return result, nil
}

func (service *BlogServiceImpl) DeleteBlog(id string) error {
	blog, err := service.store.GetBlogById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(blog.MainImage.String); err != nil {
		return err
	}

	if err := service.store.DeleteBlog(id); err != nil {
		return err
	}

	return nil
}
