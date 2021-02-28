package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TpiRepository interface {
	Create(tpi *entities.Tpi) error
}

type tpiRepository struct {
	db gorm.DB
}

func (t *tpiRepository) Create(tpi *entities.Tpi) error {
	err := t.db.Create(&tpi).Error
	return err
}

func NewTpiRepository(database gorm.DB) TpiRepository {
	return &tpiRepository{db: database}
}