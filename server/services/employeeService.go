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

type EmployeeService interface {
	GetEmployees(ctx context.Context) ([]types.Employee, error)
	CreateEmployee(ctx context.Context, employee *db.Employee) (*types.ModelResult, error)
	UpdateEmployee(ctx context.Context, id string, employee *db.Employee) (*types.ModelResult, error)
	DeleteEmployee(ctx context.Context, id string) error
}

type EmployeeServiceImpl struct {
	store    repository.EmployeeStore
	s3Client lib.S3Client
	folder   string
}

func NewEmployeeService(store repository.EmployeeStore, s3Client lib.S3Client, folder string) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{store: store, s3Client: s3Client, folder: folder}
}

func (service *EmployeeServiceImpl) GetEmployees(ctx context.Context) ([]types.Employee, error) {
	employees, err := service.store.GetEmployees(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.Employee
	for _, employee := range employees {
		result = append(result, lib.SerializeEmployee(employee))
	}
	return result, nil
}

func (service *EmployeeServiceImpl) CreateEmployee(ctx context.Context, employee *db.Employee) (*types.ModelResult, error) {
	image := employee.ProfileImage

	if !image.Valid {
		return nil, errors.New("image is empty")
	}

	PresignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, image.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	employee.ProfileImage = sql.NullString{
		String: imageUrl,
		Valid:  true,
	}

	if err := service.store.CreateEmployee(ctx, employee); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *EmployeeServiceImpl) UpdateEmployee(ctx context.Context, id string, employee *db.Employee) (*types.ModelResult, error) {
	image := employee.ProfileImage

	var PresignedUrl, imageUrl string
	var err error

	if image.Valid {
		PresignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, image.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		employee.ProfileImage = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}

	if err := service.store.UpdateEmployee(ctx, id, employee); err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: PresignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, id string) error {
	employee, err := service.store.GetEmployeeById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(employee.ProfileImage.String); err != nil {
		return err
	}

	if err := service.store.DeleteEmployee(ctx, id); err != nil {
		return err
	}

	return nil
}
