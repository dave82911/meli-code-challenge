package resources

import (
	"encoding/json"
	"meli-code-challenge/model"
	"net/http"
	"strconv"
)

const getItemDetail = "https://api.mercadolibre.com/items/"
const getCategoryDetail = "https://api.mercadolibre.com/categories/"
const getCurrencyDetail = "https://api.mercadolibre.com/currencies/"
const getUserDetail = "https://api.mercadolibre.com/users/"

type MeliResourceImpl struct {
}

func NewMeliResourceImpl() MeliResource {
	return &MeliResourceImpl{}
}

func (i *MeliResourceImpl) GetItem(site string, itemId string) (*model.ItemApiResponse, error) {
	// Make GET request to the external API
	resp, err := http.Get(getItemDetail + site + itemId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse model.ItemApiResponse
	// Check if the response status code is not in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}

	// Decode the response body into the struct representing the API response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func (i *MeliResourceImpl) GetCategoryDetail(categoryId string) (*model.CategoryApiResponse, error) {
	// Make GET request to the external API
	resp, err := http.Get(getCategoryDetail + categoryId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse model.CategoryApiResponse
	// Check if the response status code is not in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}

	// Decode the response body into the struct representing the API response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func (i *MeliResourceImpl) GetCurrencyDetail(currency string) (*model.CurrencyApiResponse, error) {
	// Make GET request to the external API
	resp, err := http.Get(getCurrencyDetail + currency)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse model.CurrencyApiResponse
	// Check if the response status code is not in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}

	// Decode the response body into the struct representing the API response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func (i *MeliResourceImpl) GetUserDetail(sellerId int) (*model.UserApiResponse, error) {
	// Make GET request to the external API
	resp, err := http.Get(getUserDetail + strconv.Itoa(sellerId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse model.UserApiResponse
	// Check if the response status code is not in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, err
	}

	// Decode the response body into the struct representing the API response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
