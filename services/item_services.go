package services

import (
	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/rest_items-api/domain/items"
)

var (
	// ItemService var from itemsServiceInterface interface
	ItemService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	Get(string) (*items.Item, errors.ResError)
	Create(items.Item) (*items.Item, errors.ResError)
}

type itemService struct{}

func (s *itemService) Get(id string) (*items.Item, errors.ResError) {
	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Create(item items.Item) (*items.Item, errors.ResError) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil

}
