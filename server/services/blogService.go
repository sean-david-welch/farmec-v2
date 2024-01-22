package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type BlogService interface {
	GetBlogs() ([]types.Blog, error)
	GetBlogsByID(id string) (*types.Blog, error)
	CreateBlog(blog *types.Blog) (*types.ModelResult, error)
	UpdateBlog(id string, blog *types.Blog) (*types.ModelResult, error)
	DeleteBlog(id string) error
}

type BlogServiceImpl struct {
	repository repository.BlogRepository
	s3Client   utils.S3Client
	folder     string
}

func NewBlogService(repository repository.BlogRepository, s3Client utils.S3Client, folder string) *BlogServiceImpl {
	return &BlogServiceImpl{repository: repository, s3Client: s3Client, folder: folder}
}

func (service *BlogServiceImpl) GetBlogs() ([]types.Blog, error) {
	blogs, err := service.repository.GetBlogs()
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (service *BlogServiceImpl) GetBlogsByID(id string) (*types.Blog, error) {
	blog, err := service.repository.GetBlogById(id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (service *BlogServiceImpl) CreateBlog(blog *types.Blog) (*types.ModelResult, error) {
	image := blog.MainImage

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image)
	if err != nil {
		return nil, err
	}

	blog.MainImage = imageUrl

	if err := service.repository.CreateBlog(blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *BlogServiceImpl) UpdateBlog(id string, blog *types.Blog) (*types.ModelResult, error) {
	image := blog.MainImage

	var presignedUrl, imageUrl string
	var err error

	if image == "" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image)
		if err != nil {
			return nil, err
		}
		blog.MainImage = imageUrl
	}

	if err := service.repository.UpdateBlog(id, blog); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl:     image,
	}

	return result, nil
}

func (service *BlogServiceImpl) DeleteBlog(id string) error {
	blog, err := service.repository.GetBlogById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(blog.MainImage); err != nil {
		return err
	}

	if err := service.repository.DeleteBlog(id); err != nil {
		return err
	}

	return nil
}
