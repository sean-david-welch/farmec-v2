package services

import (
	"errors"

	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type MachineService struct {
	folder string 
	s3Client *utils.S3Client
	repository *repository.MachineRepository
}

func NewMachineService(repository *repository.MachineRepository, s3Client *utils.S3Client, folder string) *MachineService {
	return &MachineService{
		repository: repository,
		s3Client: s3Client,
		folder: folder,
	}
}

func (service *MachineService) GetMachines(id string) ([]models.Machine, error) {
	return service.repository.GetMachines(id)
}

func (service *MachineService) CreateMachine(machine *models.Machine) (*types.MachineResult, error) {
	machineImage := machine.MachineImage

	if machineImage == "" {
		return nil, errors.New("machine image is Empty")
	}

	presignedUrl, machineUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, machineImage)
	if err != nil {
		return nil, err
	}

	machine.MachineImage = machineUrl

	service.repository.CreateMachine(machine); if err != nil {
		return nil, err
	}

	result := &types.MachineResult{
		PresginedMachine: presignedUrl,
		MachineUrl: machineUrl,
	}

	return result, nil
}

func(service *MachineService) UpdateMachine(id string, machine *models.Machine) (*types.MachineResult, error) {
	machineImage := machine.MachineImage

	var presignedUrl, machineUrl string
	var err error

	if machineImage != "" {
		presignedUrl, machineUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, machineImage)
		if err != nil {
			return nil, err
		}
		machine.MachineImage = machineUrl
	}

	service.repository.UpdateMachine(id, machine); if err != nil {
		return nil, err
	}

	result := &types.MachineResult{
		PresginedMachine: presignedUrl,
		MachineUrl: machineUrl,
	}

	return result, nil
}

func (service *MachineService) DeleteMachine(id string) error {
	machine, err := service.repository.GetMachineById(id); if err != nil {
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