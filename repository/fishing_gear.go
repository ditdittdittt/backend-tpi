package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishingGearRepository interface {
	Create(fishingGear *entities.FishingGear) error
}

type fishingGearRepository struct {
	db gorm.DB
}

func (f *fishingGearRepository) Create(fishingGear *entities.FishingGear) error {
	err := f.db.Create(&fishingGear).Error
	return err
}

func NewFishingGearRepository(database gorm.DB) FishingGearRepository {
	return &fishingGearRepository{db: database}
}