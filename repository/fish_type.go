package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishTypeRepository interface {
	Create(fishType *entities.FishType) error
}

type fishTypeRepository struct {
	db gorm.DB
}

func (f *fishTypeRepository) Create(fishType *entities.FishType) error {
	err := f.db.Create(&fishType).Error
	return err
}

func NewFishTypeRepository(database gorm.DB) FishTypeRepository {
	return &fishTypeRepository{db: database}
}