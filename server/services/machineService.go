package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type MachineService interface {
	GetMachines(ctx context.Context, id string) ([]types.Machine, error)
	GetMachineById(ctx context.Context, id string) (*types.Machine, error)
	CreateMachine(ctx context.Context, machine *db.Machine) (*types.ModelResult, error)
	UpdateMachine(ctx context.Context, id string, machine *db.Machine) (*types.ModelResult, error)
	DeleteMachine(ctx context.Context, id string) error
}

type MachineServiceImpl struct {
	folder   string
	s3Client lib.S3Client
	store    stores.MachineStore
}

func NewMachineService(store stores.MachineStore, s3Client lib.S3Client, folder string) *MachineServiceImpl {
	return &MachineServiceImpl{
		store:    store,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *MachineServiceImpl) GetMachines(ctx context.Context, id string) ([]types.Machine, error) {
	machines, err := service.store.GetMachines(ctx, id)
	if err != nil {
		return nil, errors.New("machines with supplier_id not foud")
	}
	var result []types.Machine
	for _, machine := range machines {
		result = append(result, lib.SerializeMachine(machine))
	}
	return result, nil
}

func (service *MachineServiceImpl) GetMachineById(ctx context.Context, id string) (*types.Machine, error) {
	machine, err := service.store.GetMachineById(ctx, id)
	if err != nil {
		return nil, errors.New("machine with id not found")
	}
	result := lib.SerializeMachine(*machine)

	return &result, nil
}

func (service *MachineServiceImpl) CreateMachine(ctx context.Context, machine *db.Machine) (*types.ModelResult, error) {
	machineImage := machine.MachineImage
	if !machineImage.Valid {
		return nil, errors.New("machine image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, machineImage.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	machine.MachineImage = sql.NullString{
		String: imageUrl,
		Valid:  true,
	}

	err = service.store.CreateMachine(ctx, machine)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *MachineServiceImpl) UpdateMachine(ctx context.Context, id string, machine *db.Machine) (*types.ModelResult, error) {
	machineImage := machine.MachineImage

	var presignedUrl, imageUrl string
	var err error

	if machineImage.Valid {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, machineImage.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		machine.MachineImage = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}

	err = service.store.UpdateMachine(ctx, id, machine)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *MachineServiceImpl) DeleteMachine(ctx context.Context, id string) error {
	machine, err := service.store.GetMachineById(ctx, id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(machine.MachineImage.String); err != nil {
		return err
	}

	if err := service.store.DeleteMachine(ctx, id); err != nil {
		return err
	}

	return nil
}
