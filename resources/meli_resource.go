package resources

import (
	"meli-code-challenge/model"
)

type MeliResource interface {
	GetItem(site string, itemId string) (*model.ItemApiResponse, error)
	GetCategoryDetail(categoryId string) (*model.CategoryApiResponse, error)
	GetCurrencyDetail(categoryId string) (*model.CurrencyApiResponse, error)
	GetUserDetail(sellerId int) (*model.UserApiResponse, error)
}
