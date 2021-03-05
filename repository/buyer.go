package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type BuyerRepository interface {
	Create(buyer *entities.Buyer) error
	GetWithSelectedField(selectedField []string) (buyers []entities.Buyer, err error)
}

type buyerRepository struct {
	db gorm.DB
}

func (b *buyerRepository) GetWithSelectedField(selectedField []string) (buyers []entities.Buyer, err error) {
	err = b.db.Select(selectedField).Find(&buyers).Error
	if err != nil {
		return nil, err
	}
	return buyers, err
}

func (b *buyerRepository) Create(buyer *entities.Buyer) error {
	err := b.db.Create(&buyer).Error
	return err
}

func NewBuyerRepository(database gorm.DB) BuyerRepository {
	return &buyerRepository{db: database}
}
