package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type EmployeeService interface {
	GetEmployees() ([]types.Employee, error) 
	CreateEmployee(employee *types.Employee) (*types.ModelResult, error) 
	UpdateEmployee(id string, employee *types.Employee) (*types.ModelResult, error) 
	DeleteEmployee(id string) error 
}

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
	s3Client *utils.S3Client
	folder string
}

func NewEmployeeService(repository repository.EmployeeRepository, s3Client *utils.S3Client, folder string) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{repository: repository, s3Client: s3Client, folder: folder}
}

func(service *EmployeeServiceImpl) GetEmployees() ([]types.Employee, error) {
	employees, err := service.repository.GetEmployees(); if err != nil {
		return nil, err
	}
	
	return employees, nil
}

func(service *EmployeeServiceImpl) CreateEmployee(employee *types.Employee) (*types.ModelResult, error) {
	image := employee.ProfileImage

	presginedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image); if err != nil {
		return nil, err
	}

	employee.ProfileImage = imageUrl

	if err := service.repository.CreateEmployee(employee); err != nil {
		return nil, err
	}

	result := &types.ModelResult {
		PresginedUrl: presginedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}

func(service *EmployeeServiceImpl) UpdateEmployee(id string, employee *types.Employee) (*types.ModelResult, error) {
	image := employee.ProfileImage

	var presginedUrl, imageUrl string
	var err error

	if image != "" {
		presginedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image); if err != nil {
			return nil, err
		}
		employee.ProfileImage = imageUrl
	}

	if err := service.repository.UpdateEmployee(id, employee); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presginedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}

func(service *EmployeeServiceImpl) DeleteEmployee(id string) error {
	employee, err := service.repository.GetEmployeeById(id); if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(employee.ProfileImage); err != nil {
		return err
	}

	if err := service.repository.DeleteEmployee(id); err != nil {
		return err
	}

	return nil
}