package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishTypeRepository interface {
	Create(fishType *entities.FishType) error
	GetWithSelectedField(selectedField []string) (fishTypes []entities.FishType, err error)
}

type fishTypeRepository struct {
	db gorm.DB
}

func (f *fishTypeRepository) GetWithSelectedField(selectedField []string) (fishTypes []entities.FishType, err error) {
	err = f.db.Select(selectedField).Find(&fishTypes).Error
	if err != nil {
		return nil, err
	}
	return fishTypes, err
}

func (f *fishTypeRepository) Create(fishType *entities.FishType) error {
	err := f.db.Create(&fishType).Error
	return err
}

func NewFishTypeRepository(database gorm.DB) FishTypeRepository {
	return &fishTypeRepository{db: database}
}
