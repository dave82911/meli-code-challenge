package service

import (
	"mime/multipart"
)

type ItemsService interface {
	ProcessItemsFile(file *multipart.FileHeader) error
}
