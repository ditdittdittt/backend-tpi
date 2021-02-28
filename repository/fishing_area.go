package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishingAreaRepository interface {
	Create(fishingArea *entities.FishingArea) error
}

type fishingAreaRepository struct {
	db gorm.DB
}

func (f *fishingAreaRepository) Create(fishingArea *entities.FishingArea) error {
	err := f.db.Create(&fishingArea).Error
	return err
}

func NewFishingAreaRepository(database gorm.DB) FishingAreaRepository {
	return &fishingAreaRepository{db: database}
}