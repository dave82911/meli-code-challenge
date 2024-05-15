package model

import "time"

type ItemApiResponse struct {
	ID         string    `json:"id"`
	Price      float32   `json:"price"`
	StartTime  time.Time `json:"start_time"`
	CategoryId string    `json:"category_id"`
	CurrencyId string    `json:"currency_id"`
	SellerId   int       `json:"seller_id"`
}

type CategoryApiResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CurrencyApiResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type UserApiResponse struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}
