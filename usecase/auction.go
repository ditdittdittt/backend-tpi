package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type AuctionUsecase interface {
	Create(auction *entities.Auction) error
	Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.Auction, error)
	Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Auction, error)
}

type auctionUsecase struct {
	auctionRepository mysql.AuctionRepository
	caughtRepository  mysql.CaughtRepository
}

func (a *auctionUsecase) Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.Auction, error) {
	queryMap := map[string]interface{}{
		"caughts.tpi_id": tpiID,
	}

	if fisherID != 0 {
		queryMap["caughts.fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["caughts.fish_type_id"] = fishTypeID
	}

	if caughtStatusID != 0 {
		queryMap["caughts.caught_status_id"] = caughtStatusID
	}

	startDate := time.Now().Format("2006-01-02")
	toDate := time.Now().String()

	auctions, err := a.auctionRepository.Get(queryMap, startDate, toDate)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Caught repository error")
	}

	return auctions, nil
}

func (a *auctionUsecase) Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Auction, error) {
	queryMap := map[string]interface{}{
		"caughts.tpi_id": tpiID,
	}

	if fisherID != 0 {
		queryMap["caughts.fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["caughts.fish_type_id"] = fishTypeID
	}

	auctions, err := a.auctionRepository.Search(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Search] Caught repository")
	}

	return auctions, nil
}

func (a *auctionUsecase) Create(auction *entities.Auction) error {

	auction.CreatedAt = time.Now()
	auction.UpdatedAt = time.Now()

	err := a.auctionRepository.Create(auction)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Auction repository error")
	}

	caught, err := a.caughtRepository.GetByID(auction.CaughtID)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Caught repository error")
	}

	updateStatus := map[string]interface{}{
		"caught_status_id": 2,
	}

	err = a.caughtRepository.Update(&caught, updateStatus)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Caught repository error")
	}

	return nil
}

func NewAuctionUsecase(auctionRepository mysql.AuctionRepository, caughtRepository mysql.CaughtRepository) AuctionUsecase {
	return &auctionUsecase{auctionRepository: auctionRepository, caughtRepository: caughtRepository}
}
