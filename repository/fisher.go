package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FisherRepository interface {
	GetWithSelectedField(selectedField []string) (fishers []entities.Fisher, err error)
	Create(fisher *entities.Fisher) error
}

type fisherRepository struct {
	db gorm.DB
}

func (f *fisherRepository) GetWithSelectedField(selectedField []string) (fishers []entities.Fisher, err error) {
	err = f.db.Select(selectedField).Find(&fishers).Error
	if err != nil {
		return nil, err
	}
	return fishers, err
}

func (f *fisherRepository) Create(fisher *entities.Fisher) error {
	err := f.db.Create(&fisher).Error
	return err
}

func NewFisherRepository(database gorm.DB) FisherRepository {
	return &fisherRepository{db: database}
}
