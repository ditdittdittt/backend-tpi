package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type TransactionUsecase interface {
	Create(transaction *entities.Transaction, auctionIDs []int) error
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
	auctionRepository     repository.AuctionRepository
	caughtRepository      repository.CaughtRepository
}

func (t *transactionUsecase) Create(transaction *entities.Transaction, auctionIDs []int) error {
	caughtIDs := make([]int, 0)

	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	for _, auctionID := range auctionIDs {
		transaction.TransactionItem = append(transaction.TransactionItem, entities.TransactionItem{
			AuctionID: auctionID,
		})
		auction, err := t.auctionRepository.GetByID(auctionID)
		if err != nil {
			return stacktrace.Propagate(err, "[GetByID] Auction repository error")
		}
		caughtIDs = append(caughtIDs, auction.CaughtID)
	}

	err := t.transactionRepository.Create(transaction)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Transaction repository error")
	}

	updateStatus := map[string]interface{}{
		"caught_status_id": 3,
	}

	err = t.caughtRepository.BulkUpdate(caughtIDs, updateStatus)
	if err != nil {
		return stacktrace.Propagate(err, "[BulkUpdate] Caught repository error")
	}

	return nil
}

func NewTransactionUsecase(
	transactionRepository repository.TransactionRepository,
	auctionRepository repository.AuctionRepository,
	caughtRepository repository.CaughtRepository,
) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
		auctionRepository:     auctionRepository,
		caughtRepository:      caughtRepository}
}
