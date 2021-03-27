package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FisherTpiRepository interface {
	Create(fisherTpi *entities.FisherTpi) error
	Index(query map[string]interface{}) ([]entities.FisherTpi, error)
}

type fisherTpiRepository struct {
	db gorm.DB
}

func (f *fisherTpiRepository) Index(query map[string]interface{}) ([]entities.FisherTpi, error) {
	var result []entities.FisherTpi

	err := f.db.Where(query).Preload("Fisher").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (f *fisherTpiRepository) Create(fisherTpi *entities.FisherTpi) error {
	err := f.db.Create(fisherTpi).Error
	if err != nil {
		return err
	}

	return nil
}

func NewFisherTpiRepository(database gorm.DB) FisherTpiRepository {
	return &fisherTpiRepository{db: database}
}
