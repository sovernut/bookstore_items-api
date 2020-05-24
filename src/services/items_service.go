package services

import (
	"net/http"

	"github.com/sovernut/bookstore_items-api/src/domain/items"
	"github.com/sovernut/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	// if err := item.Save(); err != nil {
	// 	return nil, err
	// }
	// return &item, nil
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implement", nil)
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	// item := items.Item{Id: id}

	// if err := item.Get(); err != nil {
	// 	return nil, err
	// }
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implement", nil)

}
