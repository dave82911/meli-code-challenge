package model

import "time"

type Item struct {
	Id           int       `gorm:"type:int;primary_key"`
	Site         string    `gorm:"type:varchar(10)"`
	ItemId       string    `gorm:"type:string"`
	Price        float32   `gorm:"type:float"`
	StartTime    time.Time `gorm:"type:time"`
	CategoryName string    `gorm:"type:string"`
	Description  string    `gorm:"type:string"`
	Nickname     string    `gorm:"type:string"`
	CreatedAt    time.Time `gorm:"type:time"`
	UpdatedAt    time.Time `gorm:"type:time"`
}
