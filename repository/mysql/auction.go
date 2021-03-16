package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type AuctionRepository interface {
	Create(auction *entities.Auction) error
	GetByID(id int) (auction entities.Auction, err error)
	Search(query map[string]interface{}) (auctions []entities.Auction, err error)
	Get(query map[string]interface{}, startDate string, toDate string) (auctions []entities.Auction, err error)
	Delete(id int) error
	Update(auction *entities.Auction) error
	GetPriceTotal(fishTypeID int, tpiID int, from string, to string) (float64, error)
	GetTransactionSpeed(fishTypeID int, tpiID int, from string, to string) (float64, error)
}

type auctionRepository struct {
	db gorm.DB
}

func (a *auctionRepository) GetTransactionSpeed(fishTypeID int, tpiID int, from string, to string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(AVG(
		UNIX_TIMESTAMP(a.created_at)-UNIX_TIMESTAMP(c.created_at)
	), 0) AS result
	FROM auctions AS a
	INNER JOIN caughts AS c ON a.caught_id = c.id
	WHERE a.created_at BETWEEN "%s" AND "%s" AND c.caught_status_id = 3`

	query = fmt.Sprintf(query, from, to)

	if fishTypeID != 0 {
		query = query + " AND c.fish_type_id = " + strconv.Itoa(fishTypeID)
	}

	if tpiID != 0 {
		query = query + " AND c.tpi_id = " + strconv.Itoa(tpiID)
	}

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *auctionRepository) GetPriceTotal(fishTypeID int, tpiID int, from string, to string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(
		SUM(a.price), 0) 
		FROM auctions AS a
		INNER JOIN caughts AS c ON a.caught_id = c.id
		WHERE a.created_at BETWEEN "%s" AND "%s" AND c.caught_status_id = 3`

	query = fmt.Sprintf(query, from, to)

	if fishTypeID != 0 {
		query = query + " AND c.fish_type_id = " + strconv.Itoa(fishTypeID)
	}

	if tpiID != 0 {
		query = query + " AND c.tpi_id = " + strconv.Itoa(tpiID)
	}

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *auctionRepository) Update(auction *entities.Auction) error {
	err := a.db.Model(&auction).Updates(auction).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *auctionRepository) Delete(id int) error {
	err := a.db.Delete(&entities.Auction{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *auctionRepository) Get(query map[string]interface{}, startDate string, toDate string) (auctions []entities.Auction, err error) {
	err = a.db.Where("created_at BETWEEN ? AND ?", startDate, toDate).
		Where("caught_id IN (?)", a.db.Table("caughts").Select("id").Where(query)).
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
	err = a.db.Where("caught_id IN (?)", a.db.Table("caughts").Select("id").Where(query)).
		Preload("Caught").
		Preload("Caught.Fisher").
		Preload("Caught.FishType").
		Find(&auctions).Error
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
