package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type EmployeeService interface {
	GetEmployees() ([]types.Employee, error)
	CreateEmployee(employee *types.Employee) (*types.ModelResult, error)
	UpdateEmployee(id string, employee *types.Employee) (*types.ModelResult, error)
	DeleteEmployee(id string) error
}

type EmployeeServiceImpl struct {
	store    store.EmployeeStore
	s3Client lib.S3Client
	folder   string
}

func NewEmployeeService(store store.EmployeeStore, s3Client lib.S3Client, folder string) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{store: store, s3Client: s3Client, folder: folder}
}

func (service *EmployeeServiceImpl) GetEmployees() ([]types.Employee, error) {
	employees, err := service.store.GetEmployees()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (service *EmployeeServiceImpl) CreateEmployee(employee *types.Employee) (*types.ModelResult, error) {
	image := employee.ProfileImage

	if image == "" || image == "null" {
		return nil, errors.New("image is empty")
	}

	PresignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	employee.ProfileImage = imageUrl

	if err := service.store.CreateEmployee(employee); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *EmployeeServiceImpl) UpdateEmployee(id string, employee *types.Employee) (*types.ModelResult, error) {
	image := employee.ProfileImage

	var PresignedUrl, imageUrl string
	var err error

	if image != "" && image != "null" {
		PresignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		employee.ProfileImage = imageUrl
	}

	if err := service.store.UpdateEmployee(id, employee); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *EmployeeServiceImpl) DeleteEmployee(id string) error {
	employee, err := service.store.GetEmployeeById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(employee.ProfileImage); err != nil {
		return err
	}

	if err := service.store.DeleteEmployee(id); err != nil {
		return err
	}

	return nil
}
