package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishingGearRepository interface {
	Create(fishingGear *entities.FishingGear) error
	GetWithSelectedField(selectedField []string) (fishingGears []entities.FishingGear, err error)
}

type fishingGearRepository struct {
	db gorm.DB
}

func (f *fishingGearRepository) GetWithSelectedField(selectedField []string) (fishingGears []entities.FishingGear, err error) {
	err = f.db.Select(selectedField).Find(&fishingGears).Error
	if err != nil {
		return nil, err
	}
	return fishingGears, err
}

func (f *fishingGearRepository) Create(fishingGear *entities.FishingGear) error {
	err := f.db.Create(&fishingGear).Error
	return err
}

func NewFishingGearRepository(database gorm.DB) FishingGearRepository {
	return &fishingGearRepository{db: database}
}
