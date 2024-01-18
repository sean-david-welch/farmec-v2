package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemService struct {
	repository *repository.LineItemRepository
}

func NewLineItemService(repository *repository.LineItemRepository) *LineItemService {
	return &LineItemService{repository: repository}
}

func(service *LineItemService) GetLineItems() ([]types.LineItem, error) {
	lineItems, err := service.repository.GetLineItems(); if err != nil {
		return nil, err
	}

	return lineItems, nil
}

func(service *LineItemService) GetLineItemById(id string) (*types.LineItem, error) {
	lineItem, err := service.repository.GetLineItemById(id); if err != nil {
		return nil, err
	}

	return lineItem, nil
}

func(service *LineItemService) CreateLineItem(lineItem *types.LineItem) error {
	if err := service.repository.CreateLineItem(lineItem); err != nil {
		return err
	}

	return nil
}

func(service *LineItemService) UpdateLineItem(id string, lineItem *types.LineItem) error {
	if err := service.repository.UpdateLineItem(id, lineItem); err != nil {
		return err
	}

	return nil
}

func (service *LineItemService) DeleteLineItem(id string) error {
	if err := service.repository.DeleteLineItem(id); err != nil {
		return err
	}

	return nil
}
