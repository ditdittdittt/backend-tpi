package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type BuyerRepository interface {
	Create(buyer *entities.Buyer) error
	GetWithSelectedField(selectedField []string) (buyers []entities.Buyer, err error)
	GetByID(id int) (buyer entities.Buyer, err error)
	Update(buyer *entities.Buyer) error
	Delete(id int) error
}

type buyerRepository struct {
	db gorm.DB
}

func (b *buyerRepository) GetByID(id int) (buyer entities.Buyer, err error) {
	err = b.db.First(&buyer, id).Error
	if err != nil {
		return entities.Buyer{}, err
	}
	return buyer, err
}

func (b *buyerRepository) Update(buyer *entities.Buyer) error {
	err := b.db.Model(&buyer).Updates(buyer).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *buyerRepository) Delete(id int) error {
	err := b.db.Delete(&entities.Buyer{}, id).Error
	if err != nil {
		return err
	}

	return nil
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
