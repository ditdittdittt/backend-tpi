package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	Get(query map[string]interface{}, startDate string, toDate string) (transactions []entities.Transaction, err error)
	GetByID(id int) (transaction entities.Transaction, err error)
	Update(transaction *entities.Transaction) error
	Delete(id int) error
	GetTransactionTotal(tpiID int, districtID int, from string, to string) (int, error)
	GetBuyerTotal(status string, tpiID int, districtID int, from string, to string) (int, error)

	// Dashboard
	GetBuyerTotalDashboard(tpiID int, districtID int) ([]map[string]interface{}, error)
}

type transactionRepository struct {
	db gorm.DB
}

func (t *transactionRepository) GetBuyerTotalDashboard(tpiID int, districtID int) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := `SELECT b.status AS status, COALESCE(COUNT(DISTINCT t.buyer_id), 0) AS total
		FROM transactions AS t
		INNER JOIN buyers AS b ON t.buyer_id = b.id`

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS tpi ON t.tpi_id = tpi.id WHERE tpi.district_id = " + strconv.Itoa(districtID)
	}
	query = query + " GROUP BY b.status"

	err := t.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *transactionRepository) GetBuyerTotal(status string, tpiID int, districtID int, from string, to string) (int, error) {
	var result int
	query := `SELECT COALESCE(COUNT(DISTINCT t.buyer_id), 0) 
		FROM transactions AS t 
		INNER JOIN buyers AS b ON t.buyer_id = b.id`

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS tpi ON t.tpi_id = tpi.id WHERE tpi.district_id = " + strconv.Itoa(districtID)
	}

	query = query + ` AND t.created_at BETWEEN "%s" AND "%s" AND b.status = "%s"`
	query = fmt.Sprintf(query, from, to, status)

	err := t.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (t *transactionRepository) GetTransactionTotal(tpiID int, districtID int, from string, to string) (int, error) {
	var result int

	query := `SELECT COALESCE(COUNT(*), 0) FROM transactions AS t`

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS tpi ON t.tpi_id = tpi.id WHERE tpi.district_id = " + strconv.Itoa(districtID)
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
