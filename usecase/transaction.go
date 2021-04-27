package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
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
	tpiRepository             mysql.TpiRepository
}

func (t *transactionUsecase) GetByID(id int) (entities.Transaction, error) {
	transaction, err := t.transactionRepository.GetByID(id)
	if err != nil {
		return entities.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionUsecase) Update(transaction *entities.Transaction) error {
	// insert log
	err := helper.InsertLog(transaction.ID, constant.Transaction)
	if err != nil {
		return err
	}

	err = t.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionUsecase) Delete(id int) error {
	transaction, err := t.transactionRepository.GetByID(id)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"caught_status_id": 2,
	}

	for _, transactionItem := range transaction.TransactionItem {
		err := t.caughtItemRepository.Update(transactionItem.Auction.CaughtItemID, data)
		if err != nil {
			return err
		}

		err = t.transactionItemRepository.Delete(transactionItem.ID)
		if err != nil {
			return err
		}
	}

	err = t.transactionRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionUsecase) Index(tpiID int) ([]entities.Transaction, error) {
	queryMap := map[string]interface{}{
		"transactions.tpi_id": tpiID,
	}

	date := time.Now().Format("2006-01-02")

	transactions, err := t.transactionRepository.Index(queryMap, date)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionUsecase) Create(transaction *entities.Transaction, auctionIDs []int) error {
	caughtItemsID := make([]int, 0)

	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	currentDate := time.Now().Format("2006-01-02")
	existingCode, err := t.transactionRepository.GetLatestCode(currentDate)
	if err != nil {
		return err
	}

	tpi, err := t.tpiRepository.GetByID(transaction.TpiID)
	if err != nil {
		return err
	}

	if existingCode != "" {
		latestID := existingCode[len(existingCode)-3:]
		intLatestID, _ := strconv.Atoi(latestID)
		intLatestID++
		transaction.Code = t.formatCode(currentDate) + tpi.Code + fmt.Sprintf("/%03d", intLatestID)
	} else {
		transaction.Code = t.formatCode(currentDate) + tpi.Code + "/001"
	}

	for _, auctionID := range auctionIDs {
		transaction.TransactionItem = append(transaction.TransactionItem, &entities.TransactionItem{
			AuctionID: auctionID,
		})
		auction, err := t.auctionRepository.GetByID(auctionID)
		if err != nil {
			return err
		}
		transaction.TotalPrice += auction.Price
		caughtItemsID = append(caughtItemsID, auction.CaughtItemID)
	}

	err = t.transactionRepository.Create(transaction)
	if err != nil {
		return err
	}

	updateStatus := map[string]interface{}{
		"caught_status_id": 3,
	}

	err = t.caughtItemRepository.BulkUpdate(caughtItemsID, updateStatus)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionUsecase(
	transactionRepository mysql.TransactionRepository,
	auctionRepository mysql.AuctionRepository,
	caughtRepository mysql.CaughtRepository,
	transactionItemRepository mysql.TransactionItemRepository,
	caughtItemRepository mysql.CaughtItemRepository,
	tpiRepository mysql.TpiRepository,
) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository:     transactionRepository,
		auctionRepository:         auctionRepository,
		caughtRepository:          caughtRepository,
		transactionItemRepository: transactionItemRepository,
		caughtItemRepository:      caughtItemRepository,
		tpiRepository:             tpiRepository}
}

func (t *transactionUsecase) formatCode(date string) string {
	result := strings.ReplaceAll(date, "-", "")
	return "INV/" + result[2:] + "/"
}
