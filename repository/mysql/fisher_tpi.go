package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FisherTpiRepository interface {
	Create(fisherTpi *entities.FisherTpi) error
	Index(query map[string]interface{}) ([]entities.FisherTpi, error)
	Get(query map[string]interface{}) (entities.FisherTpi, error)
	Delete(query map[string]interface{}) error
}

type fisherTpiRepository struct {
	db gorm.DB
}

func (f *fisherTpiRepository) Get(query map[string]interface{}) (entities.FisherTpi, error) {
	var result entities.FisherTpi

	err := f.db.Where(query).Preload("Fisher").First(&result).Error
	if err != nil {
		return entities.FisherTpi{}, err
	}

	return result, nil
}

func (f *fisherTpiRepository) Delete(query map[string]interface{}) error {
	err := f.db.Where(query).Delete(&entities.FisherTpi{}).Error
	if err != nil {
		return err
	}

	return nil
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
