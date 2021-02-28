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
}

func (t *transactionUsecase) Create(transaction *entities.Transaction, auctionIDs []int) error {
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	for _, auctionID := range auctionIDs {
		transaction.TransactionItem = append(transaction.TransactionItem, entities.TransactionItem{
			AuctionID: auctionID,
		})
	}

	err := t.transactionRepository.Create(transaction)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Transaction repository error")
	}

	return nil
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{transactionRepository: transactionRepository}
}
