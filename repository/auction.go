package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type AuctionRepository interface {
	Create(auction *entities.Auction) error
	GetByID(id int) (auction entities.Auction, err error)
}

type auctionRepository struct {
	db gorm.DB
}

func (a *auctionRepository) GetByID(id int) (auction entities.Auction, err error) {
	err = a.db.First(&auction, id).Error
	if err != nil {
		return entities.Auction{}, err
	}
	return auction, nil
}

func (a *auctionRepository) Create(auction *entities.Auction) error {
	err := a.db.Create(&auction).Error
	return err
}

func NewAuctionRepository(database gorm.DB) AuctionRepository {
	return &auctionRepository{db: database}
}
