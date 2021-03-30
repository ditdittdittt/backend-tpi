package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type AuctionUsecase interface {
	Create(auction *entities.Auction) error
	Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.Auction, error)
	Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Auction, error)
	GetByID(id int) (entities.Auction, error)
	Delete(id int) error
	Update(auction *entities.Auction) error
}

type auctionUsecase struct {
	auctionRepository    mysql.AuctionRepository
	caughtRepository     mysql.CaughtRepository
	caughtItemRepository mysql.CaughtItemRepository
}

func (a *auctionUsecase) Update(auction *entities.Auction) error {
	// insert log
	err := helper.InsertLog(auction.ID, constant.Auction)
	if err != nil {
		return err
	}

	err = a.auctionRepository.Update(auction)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Auction repository error")
	}

	return nil
}

func (a *auctionUsecase) Delete(id int) error {
	auction, err := a.auctionRepository.GetByID(id)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Auction repository error for auction id %d", id)
	}

	data := map[string]interface{}{
		"caught_status_id": 1,
	}

	err = a.caughtItemRepository.Update(auction.CaughtItemID, data)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Caught repository error for caught id")
	}

	err = a.auctionRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Auction repository error")
	}

	return nil
}

func (a *auctionUsecase) GetByID(id int) (entities.Auction, error) {
	auction, err := a.auctionRepository.GetByID(id)
	if err != nil {
		return entities.Auction{}, stacktrace.Propagate(err, "[GetByID] Auction repository error")
	}

	return auction, nil
}

func (a *auctionUsecase) Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.Auction, error) {
	queryMap := map[string]interface{}{
		"auctions.tpi_id": tpiID,
	}

	if fisherID != 0 {
		queryMap["caughts.fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["caught_items.fish_type_id"] = fishTypeID
	}

	if caughtStatusID != 0 {
		queryMap["caught_items.caught_status_id"] = caughtStatusID
	}

	date := time.Now().Format("2006-01-02")

	auctions, err := a.auctionRepository.Index(queryMap, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Caught repository error")
	}

	return auctions, nil
}

func (a *auctionUsecase) Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Auction, error) {
	queryMap := map[string]interface{}{
		"auctions.tpi_id":               tpiID,
		"caught_items.caught_status_id": 2,
	}

	if fisherID != 0 {
		queryMap["caughts.fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["caught_items.fish_type_id"] = fishTypeID
	}

	auctions, err := a.auctionRepository.Search(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Search] Auction repository error")
	}

	return auctions, nil
}

func (a *auctionUsecase) Create(auction *entities.Auction) error {

	auction.CreatedAt = time.Now()
	auction.UpdatedAt = time.Now()

	caughtItem, err := a.caughtItemRepository.GetByID(auction.CaughtItemID)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Caught item repository error")
	}

	auction.Code = caughtItem.Code

	err = a.auctionRepository.Create(auction)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Auction repository error")
	}

	updateStatus := map[string]interface{}{
		"caught_status_id": 2,
	}

	err = a.caughtItemRepository.Update(auction.CaughtItemID, updateStatus)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Caught item repository error")
	}

	return nil
}

func NewAuctionUsecase(auctionRepository mysql.AuctionRepository, caughtRepository mysql.CaughtRepository, caughtItemRepository mysql.CaughtItemRepository) AuctionUsecase {
	return &auctionUsecase{auctionRepository: auctionRepository, caughtRepository: caughtRepository, caughtItemRepository: caughtItemRepository}
}
