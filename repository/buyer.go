package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type BuyerRepository interface {
	Create(buyer *entities.Buyer) error
}

type buyerRepository struct {
	db gorm.DB
}

func (b *buyerRepository) Create(buyer *entities.Buyer) error {
	err := b.db.Create(&buyer).Error
	return err
}

func NewBuyerRepository(database gorm.DB) BuyerRepository {
	return &buyerRepository{db: database}
}