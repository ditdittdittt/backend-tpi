package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type AuctionRepository interface {
	Create(auction *entities.Auction) error
}

type auctionRepository struct {
	db gorm.DB
}

func (a *auctionRepository) Create(auction *entities.Auction) error {
	err := a.db.Create(&auction).Error
	return err
}

func NewAuctionRepository(database gorm.DB) AuctionRepository {
	return &auctionRepository{db: database}
}