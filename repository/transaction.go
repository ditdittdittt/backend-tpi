package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) error
	Get() (transactions []entities.Transaction, err error)
}

type transactionRepository struct {
	db gorm.DB
}

func (t *transactionRepository) Get() (transactions []entities.Transaction, err error) {
	err = t.db.Preload("Buyer").Preload("TransactionItem").Preload("TransactionItem.Auction").Find(&transactions).Error
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
