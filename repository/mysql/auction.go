package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type AuctionRepository interface {
	Create(auction *entities.Auction) error
	GetByID(id int) (auction entities.Auction, err error)
	Search(query map[string]interface{}) (auctions []entities.Auction, err error)
	Get(query map[string]interface{}, startDate string, toDate string) (auctions []entities.Auction, err error)
}

type auctionRepository struct {
	db gorm.DB
}

func (a *auctionRepository) Get(query map[string]interface{}, startDate string, toDate string) (auctions []entities.Auction, err error) {
	err = a.db.Where("created_at BETWEEN ? AND ?", startDate, toDate).
		Preload("Caught", query).
		Preload("Caught").
		Preload("Caught.Fisher").
		Preload("Caught.FishType").
		Preload("Caught.CaughtStatus").
		Find(&auctions).Error
	if err != nil {
		return nil, err
	}

	return auctions, nil
}

func (a *auctionRepository) Search(query map[string]interface{}) (auctions []entities.Auction, err error) {
	err = a.db.Where("caught_id IN (?)", a.db.Table("caughts").Select("id").Where(query)).Preload("Caught").Preload("Caught.Fisher").Preload("Caught.FishType").Find(&auctions).Error
	if err != nil {
		return nil, err
	}
	return auctions, nil
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
