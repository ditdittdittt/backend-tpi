package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type BuyerTpiRepository interface {
	Create(buyerTpi *entities.BuyerTpi) error
	Index(query map[string]interface{}) ([]entities.BuyerTpi, error)
	Get(query map[string]interface{}) (entities.BuyerTpi, error)
	Delete(query map[string]interface{}) error
}

type buyerTpiRepository struct {
	db gorm.DB
}

func (b *buyerTpiRepository) Get(query map[string]interface{}) (entities.BuyerTpi, error) {
	var result entities.BuyerTpi

	err := b.db.Where(query).Preload("Buyer").First(&result).Error
	if err != nil {
		return entities.BuyerTpi{}, err
	}

	return result, nil
}

func (b *buyerTpiRepository) Delete(query map[string]interface{}) error {
	err := b.db.Where(query).Delete(&entities.BuyerTpi{}).Error
	if err != nil {
		return err
	}

	return nil
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
