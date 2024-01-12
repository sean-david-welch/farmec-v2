package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type BlogService struct {
	repository *repository.BlogRepository
	s3Client *utils.S3Client
	folder string
}

func NewBlogService(repository *repository.BlogRepository, s3Client *utils.S3Client, folder string) *BlogService {
	return &BlogService{repository: repository, s3Client: s3Client, folder: folder}
}

func(service *BlogService) GetBlogs() ([]models.Blog, error) {
	blogs, err := service.repository.GetBlogs(); if err != nil {
		return nil, err
	}

	return blogs, nil
}

func(service *BlogService) GetBlogsByID(id string) (*models.Blog, error) {
	blog, err := service.repository.GetBlogById(id); if err != nil {
		return nil, err
	}

	return blog, nil
}

func(service *BlogService) CreateBlog(blog *models.Blog) (*types.ModelResult, error) {
	image := blog.MainImage

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image); if err != nil {
		return nil, err
	}

	blog.MainImage = imageUrl

	if err := service.repository.CreateBlog(blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: image,
	}

	return result, nil
}

func(service *BlogService) UpdateBlog(id string, blog *models.Blog) (*types.ModelResult, error) {
	image := blog.MainImage

	var presignedUrl, imageUrl string
	var err error

	if image != "" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image); if err != nil {
			return nil, err
		}
		blog.MainImage = imageUrl
	}


	if err := service.repository.UpdateBlog(id, blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: image,
	}

	return result, nil
}

func(service *BlogService) DeleteBlog(id string) error {
	blog, err := service.repository.GetBlogById(id); if err != nil {
		return nil, err
	}

	return nil
}