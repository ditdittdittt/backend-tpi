package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type TransactionUsecase interface {
	Create(transaction *entities.Transaction, auctionIDs []int) error
	Index(tpiID int) ([]entities.Transaction, error)
	GetByID(id int) (entities.Transaction, error)
	Update(transaction *entities.Transaction) error
	Delete(id int) error
}

type transactionUsecase struct {
	transactionRepository     mysql.TransactionRepository
	auctionRepository         mysql.AuctionRepository
	caughtRepository          mysql.CaughtRepository
	caughtItemRepository      mysql.CaughtItemRepository
	transactionItemRepository mysql.TransactionItemRepository
}

func (t *transactionUsecase) GetByID(id int) (entities.Transaction, error) {
	transaction, err := t.transactionRepository.GetByID(id)
	if err != nil {
		return entities.Transaction{}, stacktrace.Propagate(err, "[GetByID] Transaction repository error")
	}

	return transaction, nil
}

func (t *transactionUsecase) Update(transaction *entities.Transaction) error {
	err := t.transactionRepository.Update(transaction)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Transaction repository error")
	}

	return nil
}

func (t *transactionUsecase) Delete(id int) error {
	//transaction, err := t.transactionRepository.GetByID(id)
	//if err != nil {
	//	return stacktrace.Propagate(err, "[GetByID] Transaction repository error")
	//}
	//
	//data := map[string]interface{}{
	//	"caught_status_id": 2,
	//}
	//
	//for _, transactionItem := range transaction.TransactionItem {
	//	err := t.caughtRepository.Update(transactionItem.Auction.Caught, data)
	//	if err != nil {
	//		return stacktrace.Propagate(err, "[Update] Transaction repository error")
	//	}
	//
	//	err = t.transactionItemRepository.Delete(transactionItem.ID)
	//	if err != nil {
	//		return stacktrace.Propagate(err, "[Delete] Transaction item repository error")
	//	}
	//}
	//
	//err = t.transactionRepository.Delete(id)
	//if err != nil {
	//	return stacktrace.Propagate(err, "[Delete] Transaction repository error")
	//}

	return nil
}

func (t *transactionUsecase) Index(tpiID int) ([]entities.Transaction, error) {
	queryMap := map[string]interface{}{
		"transactions.tpi_id": tpiID,
	}

	date := time.Now().Format("2006-01-02")

	transactions, err := t.transactionRepository.Index(queryMap, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Transaction repository error")
	}

	return transactions, nil
}

func (t *transactionUsecase) Create(transaction *entities.Transaction, auctionIDs []int) error {
	caughtItemsID := make([]int, 0)

	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	for _, auctionID := range auctionIDs {
		transaction.TransactionItem = append(transaction.TransactionItem, &entities.TransactionItem{
			AuctionID: auctionID,
		})
		auction, err := t.auctionRepository.GetByID(auctionID)
		if err != nil {
			return stacktrace.Propagate(err, "[GetByID] Auction repository error")
		}
		transaction.TotalPrice += auction.Price
		caughtItemsID = append(caughtItemsID, auction.CaughtItemID)
	}

	err := t.transactionRepository.Create(transaction)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Transaction repository error")
	}

	updateStatus := map[string]interface{}{
		"caught_status_id": 3,
	}

	err = t.caughtItemRepository.BulkUpdate(caughtItemsID, updateStatus)
	if err != nil {
		return stacktrace.Propagate(err, "[BulkUpdate] Caught repository error")
	}

	return nil
}

func NewTransactionUsecase(
	transactionRepository mysql.TransactionRepository,
	auctionRepository mysql.AuctionRepository,
	caughtRepository mysql.CaughtRepository,
	transactionItemRepository mysql.TransactionItemRepository,
	caughtItemRepository mysql.CaughtItemRepository,
) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository:     transactionRepository,
		auctionRepository:         auctionRepository,
		caughtRepository:          caughtRepository,
		transactionItemRepository: transactionItemRepository,
		caughtItemRepository:      caughtItemRepository}
}
