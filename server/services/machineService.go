package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type MachineService interface {
	GetMachines(id string) ([]types.Machine, error)
	GetMachineById(id string) (*types.Machine, error)
	CreateMachine(machine *types.Machine) (*types.ModelResult, error)
	UpdateMachine(id string, machine *types.Machine) (*types.ModelResult, error)
	DeleteMachine(id string) error
}

type MachineServiceImpl struct {
	folder     string
	s3Client   lib.S3Client
	repository repository.MachineRepository
}

func NewMachineService(repository repository.MachineRepository, s3Client lib.S3Client, folder string) *MachineServiceImpl {
	return &MachineServiceImpl{
		repository: repository,
		s3Client:   s3Client,
		folder:     folder,
	}
}

func (service *MachineServiceImpl) GetMachines(id string) ([]types.Machine, error) {
	machines, err := service.repository.GetMachines(id)
	if err != nil {
		return nil, errors.New("machines with supplier_id not foud")
	}

	return machines, nil
}

func (service *MachineServiceImpl) GetMachineById(id string) (*types.Machine, error) {
	machine, err := service.repository.GetMachineById(id)
	if err != nil {
		return nil, errors.New("machine with id not found")
	}

	return machine, nil
}

func (service *MachineServiceImpl) CreateMachine(machine *types.Machine) (*types.ModelResult, error) {
	machineImage := machine.MachineImage
	if machineImage == "" {
		return nil, errors.New("machine image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, machineImage)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	machine.MachineImage = imageUrl

	service.repository.CreateMachine(machine)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *MachineServiceImpl) UpdateMachine(id string, machine *types.Machine) (*types.ModelResult, error) {
	machineImage := machine.MachineImage

	if machineImage == "" || machineImage == "null" {
		return nil, errors.New("image is empty")
	}

	var presignedUrl, imageUrl string
	var err error

	if machineImage != "" && machineImage != "null" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, machineImage)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		machine.MachineImage = imageUrl
	}

	service.repository.UpdateMachine(id, machine)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *MachineServiceImpl) DeleteMachine(id string) error {
	machine, err := service.repository.GetMachineById(id)
	if err != nil {
		return err
	}

	if err := service.s3Client.DeleteImageFromS3(machine.MachineImage); err != nil {
		return err
	}

	if err := service.repository.DeleteMachine(id); err != nil {
		return err
	}

	return nil
}
