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
}

func (a *auctionUsecase) Create(auction *entities.Auction) error {
	auction.CreatedAt = time.Now()
	auction.UpdatedAt = time.Now()

	err := a.auctionRepository.Create(auction)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Auction repository error")
	}

	return nil
}

func NewAuctionUsecase(auctionRepository repository.AuctionRepository) AuctionUsecase {
	return &auctionUsecase{auctionRepository: auctionRepository}
}
