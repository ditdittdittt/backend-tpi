package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FisherRepository interface {
	Create(fisher *entities.Fisher) error
}

type fisherRepository struct {
	db gorm.DB
}

func (f *fisherRepository) Create(fisher *entities.Fisher) error {
	err := f.db.Create(&fisher).Error
	return err
}

func NewFisherRepository(database gorm.DB) FisherRepository {
	return &fisherRepository{db: database}
}