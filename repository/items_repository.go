package repository

import (
	"gorm.io/gorm"
)

type ItemRepository interface {
	BeginTransaction() *gorm.DB
}
