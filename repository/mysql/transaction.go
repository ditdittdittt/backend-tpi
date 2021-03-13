package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	Get(startDate string, toDate string) (transactions []entities.Transaction, err error)
	GetByID(id int) (transaction entities.Transaction, err error)
	Update(transaction *entities.Transaction) error
	Delete(id int) error
}

type transactionRepository struct {
	db gorm.DB
}

func (t *transactionRepository) GetByID(id int) (transaction entities.Transaction, err error) {
	err = t.db.First(&transaction, id).Error
	if err != nil {
		return entities.Transaction{}, err
	}
	return transaction, nil
}

func (t *transactionRepository) Update(transaction *entities.Transaction) error {
	panic("implement me")
}

func (t *transactionRepository) Delete(id int) error {
	err := t.db.Delete(&entities.Transaction{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) Get(startDate string, toDate string) (transactions []entities.Transaction, err error) {
	err = t.db.Where("created_at BETWEEN ? AND ?", startDate, toDate).
		Preload("Buyer").
		Preload("TransactionItem").
		Preload("TransactionItem.Auction").
		Preload("TransactionItem.Auction.Caught").
		Preload("TransactionItem.Auction.Caught.Fisher").
		Preload("TransactionItem.Auction.Caught.FishType").Find(&transactions).Error
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
