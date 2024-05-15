package service

import (
	"encoding/csv"
	"io"
	"meli-code-challenge/model"
	"meli-code-challenge/repository"
	"meli-code-challenge/resources"
	"mime/multipart"
)

type ItemServiceImpl struct {
	ItemRepository repository.ItemRepository
	MeliResource   resources.MeliResource
}

func NewItemsServiceImpl(
	itemRepository repository.ItemRepository,
	meliResource resources.MeliResource,
) ItemsService {
	return &ItemServiceImpl{
		ItemRepository: itemRepository,
		MeliResource:   meliResource,
	}
}

// ProcessItemsFile Update implements ItemsService
func (t *ItemServiceImpl) ProcessItemsFile(file *multipart.FileHeader) error {
	// Open the uploaded file
	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	// Parse the CSV file
	reader := csv.NewReader(uploadedFile)
	// Read and ignore the header row
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		return err
	}

	// Start a transaction
	tx := t.ItemRepository.BeginTransaction()

	// Read each row and insert into the database
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			tx.Rollback()

			return err
		}

		// Create a new item object and insert into the database
		site := record[0]
		itemId := record[1]

		itemFromMeli, _ := t.MeliResource.GetItem(site, itemId)
		if itemFromMeli == nil {
			continue
		}

		categoryFromMeli, _ := t.MeliResource.GetCategoryDetail(itemFromMeli.CategoryId)
		if categoryFromMeli == nil {
			categoryFromMeli = &model.CategoryApiResponse{
				Name: "",
			}
		}

		currencyFromMeli, _ := t.MeliResource.GetCurrencyDetail(itemFromMeli.CurrencyId)
		if currencyFromMeli == nil {
			currencyFromMeli = &model.CurrencyApiResponse{
				Description: "",
			}
		}

		userFromMeli, _ := t.MeliResource.GetUserDetail(itemFromMeli.SellerId)
		if userFromMeli == nil {
			userFromMeli = &model.UserApiResponse{
				Nickname: "",
			}
		}

		item := model.Item{
			Site:         site,
			ItemId:       itemId,
			Price:        itemFromMeli.Price,
			StartTime:    itemFromMeli.StartTime,
			CategoryName: categoryFromMeli.Name,
			Description:  currencyFromMeli.Description,
			Nickname:     userFromMeli.Nickname,
		}

		itemFound := model.Item{}
		tx.Find(&itemFound, &model.Item{
			Site:   site,
			ItemId: itemId,
		})

		if itemFound.ItemId == "" {
			if err := tx.Create(&item).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
