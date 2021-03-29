package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	Get(query map[string]interface{}, startDate string, toDate string) (transactions []entities.Transaction, err error)
	GetByID(id int) (transaction entities.Transaction, err error)
	Update(transaction *entities.Transaction) error
	Delete(id int) error
	Index(query map[string]interface{}, date string) ([]entities.Transaction, error)

	// Report
	GetTransactionTotal(tpiID int, from string, to string) (int, error)
	GetBuyerTotal(status string, tpiID int, from string, to string) (int, error)

	// Dashboard
	GetBuyerTotalDashboard(tpiID int, status string) (int, error)
}

type transactionRepository struct {
	db gorm.DB
}

func (t *transactionRepository) Index(query map[string]interface{}, date string) ([]entities.Transaction, error) {
	var result []entities.Transaction

	err := t.db.Table("transactions").
		Where("DATE(transactions.created_at) = DATE(?)", date).
		Where(query).
		Preload("Buyer").
		Preload("TransactionItem.Auction.CaughtItem.Caught.Fisher").
		Preload("TransactionItem.Auction.CaughtItem.FishType").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, err
}

func (t *transactionRepository) GetBuyerTotalDashboard(tpiID int, status string) (int, error) {
	var result int

	query := `SELECT COALESCE(COUNT(DISTINCT t.buyer_id), 0) AS total
		FROM transactions AS t`

	switch status {
	case constant.PermanentStatus:
		query = query + " INNER JOIN buyers AS b ON t.buyer_id = b.id AND b.tpi_id = " + strconv.Itoa(tpiID)
	case constant.TemporaryStatus:
		query = query + " INNER JOIN buyer_tpis AS bt ON t.buyer_id = bt.buyer_id AND bt.tpi_id = " + strconv.Itoa(tpiID)
	}

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	err := t.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (t *transactionRepository) GetBuyerTotal(status string, tpiID int, from string, to string) (int, error) {
	var result int
	query := `SELECT COALESCE(COUNT(DISTINCT t.buyer_id), 0) 
		FROM transactions AS t`

	switch status {
	case constant.PermanentStatus:
		query = query + " INNER JOIN buyers AS b ON t.buyer_id = b.id AND b.tpi_id = " + strconv.Itoa(tpiID)
	case constant.TemporaryStatus:
		query = query + " INNER JOIN buyer_tpis AS bt ON t.buyer_id = bt.buyer_id AND bt.tpi_id = " + strconv.Itoa(tpiID)
	}

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	query = query + ` AND t.created_at BETWEEN "%s" AND "%s"`
	query = fmt.Sprintf(query, from, to)

	err := t.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (t *transactionRepository) GetTransactionTotal(tpiID int, from string, to string) (int, error) {
	var result int

	query := `SELECT COALESCE(COUNT(*), 0) FROM transactions AS t`

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	query = query + ` AND t.created_at BETWEEN "%s" AND "%s"`
	query = fmt.Sprintf(query, from, to)

	err := t.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (t *transactionRepository) GetByID(id int) (transaction entities.Transaction, err error) {
	err = t.db.Preload("TransactionItem").Preload("TransactionItem.Auction").Preload("TransactionItem.Auction.Caught").First(&transaction, id).Error
	if err != nil {
		return entities.Transaction{}, err
	}
	return transaction, nil
}

func (t *transactionRepository) Update(transaction *entities.Transaction) error {
	err := t.db.Model(&transaction).Updates(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepository) Delete(id int) error {
	err := t.db.Delete(&entities.Transaction{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) Get(query map[string]interface{}, startDate string, toDate string) (transactions []entities.Transaction, err error) {
	err = t.db.Where("created_at BETWEEN ? AND ?", startDate, toDate).
		Preload("Buyer").
		Preload("TransactionItem").
		Preload("TransactionItem.Auction").
		Preload("TransactionItem.Auction.Caught").
		Preload("TransactionItem.Auction.Caught.Fisher").
		Preload("TransactionItem.Auction.Caught.FishType").Find(&transactions, query).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionRepository) Create(transaction *entities.Transaction) error {
	err := t.db.Create(&transaction).Error
	return err
}

func NewTransactionRepository(database gorm.DB) TransactionRepository {
	return &transactionRepository{db: database}
}
