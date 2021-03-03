package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type AuctionUsecase interface {
	Create(auction *entities.Auction) error
}

type auctionUsecase struct {
	auctionRepository repository.AuctionRepository
	caughtRepository  repository.CaughtRepository
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

func NewAuctionUsecase(auctionRepository repository.AuctionRepository, caughtRepository repository.CaughtRepository) AuctionUsecase {
	return &auctionUsecase{auctionRepository: auctionRepository, caughtRepository: caughtRepository}
}
