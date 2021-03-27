package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type BuyerTpiRepository interface {
	Create(buyerTpi *entities.BuyerTpi) error
	Index(query map[string]interface{}) ([]entities.BuyerTpi, error)
}

type buyerTpiRepository struct {
	db gorm.DB
}

func (b *buyerTpiRepository) Create(buyerTpi *entities.BuyerTpi) error {
	err := b.db.Create(buyerTpi).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *buyerTpiRepository) Index(query map[string]interface{}) ([]entities.BuyerTpi, error) {
	var result []entities.BuyerTpi

	err := b.db.Where(query).Preload("Buyer").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewBuyerTpiRepository(database gorm.DB) BuyerTpiRepository {
	return &buyerTpiRepository{db: database}
}
