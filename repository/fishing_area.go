package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishingAreaRepository interface {
	Create(fishingArea *entities.FishingArea) error
	GetWithSelectedField(selectedField []string) (fishingAreas []entities.FishingArea, err error)
}

type fishingAreaRepository struct {
	db gorm.DB
}

func (f *fishingAreaRepository) GetWithSelectedField(selectedField []string) (fishingAreas []entities.FishingArea, err error) {
	err = f.db.Select(selectedField).Find(&fishingAreas).Error
	if err != nil {
		return nil, err
	}
	return fishingAreas, err
}

func (f *fishingAreaRepository) Create(fishingArea *entities.FishingArea) error {
	err := f.db.Create(&fishingArea).Error
	return err
}

func NewFishingAreaRepository(database gorm.DB) FishingAreaRepository {
	return &fishingAreaRepository{db: database}
}
