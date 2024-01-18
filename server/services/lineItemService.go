package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemService interface {
	GetLineItems() ([]types.LineItem, error)
	GetLineItemById(id string) (*types.LineItem, error)
	CreateLineItem(lineItem *types.LineItem) error
	UpdateLineItem(id string, lineItem *types.LineItem) error
	DeleteLineItem(id string) error
}

type LineItemServiceImpl struct {
	repository repository.LineItemRepository
}

func NewLineItemService(repository repository.LineItemRepository) *LineItemServiceImpl {
	return &LineItemServiceImpl{repository: repository}
}

func(service *LineItemServiceImpl) GetLineItems() ([]types.LineItem, error) {
	lineItems, err := service.repository.GetLineItems(); if err != nil {
		return nil, err
	}

	return lineItems, nil
}

func(service *LineItemServiceImpl) GetLineItemById(id string) (*types.LineItem, error) {
	lineItem, err := service.repository.GetLineItemById(id); if err != nil {
		return nil, err
	}

	return lineItem, nil
}

func(service *LineItemServiceImpl) CreateLineItem(lineItem *types.LineItem) error {
	if err := service.repository.CreateLineItem(lineItem); err != nil {
		return err
	}

	return nil
}

func(service *LineItemServiceImpl) UpdateLineItem(id string, lineItem *types.LineItem) error {
	if err := service.repository.UpdateLineItem(id, lineItem); err != nil {
		return err
	}

	return nil
}

func (service *LineItemServiceImpl) DeleteLineItem(id string) error {
	if err := service.repository.DeleteLineItem(id); err != nil {
		return err
	}

	return nil
}
