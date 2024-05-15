package repository

import (
	"gorm.io/gorm"
	"mime/multipart"
)

type ItemsRepositoryImpl struct {
	Db *gorm.DB
}

func NewItemsRepositoryImpl(Db *gorm.DB) ItemRepository {
	return &ItemsRepositoryImpl{Db: Db}
}

func (t *ItemsRepositoryImpl) ProcessItemsFile(file *multipart.FileHeader) error {
	return nil
}

func (t *ItemsRepositoryImpl) BeginTransaction() *gorm.DB {
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return tx
}
