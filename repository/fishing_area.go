package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishingAreaRepository interface {
	Create(fishingArea *entities.FishingArea) error
	GetWithSelectedField(selectedField []string) (fishingAreas []entities.FishingArea, err error)
	GetByID(id int) (fishingArea entities.FishingArea, err error)
	Update(fishingArea *entities.FishingArea) error
	Delete(id int) error
}

type fishingAreaRepository struct {
	db gorm.DB
}

func (f *fishingAreaRepository) GetByID(id int) (fishingArea entities.FishingArea, err error) {
	err = f.db.Preload("District").First(&fishingArea, id).Error
	if err != nil {
		return entities.FishingArea{}, err
	}
	return fishingArea, nil
}

func (f *fishingAreaRepository) Update(fishingArea *entities.FishingArea) error {
	err := f.db.Model(&fishingArea).Updates(fishingArea).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *fishingAreaRepository) Delete(id int) error {
	err := f.db.Delete(&entities.FishingArea{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (f *fishingAreaRepository) GetWithSelectedField(selectedField []string) (fishingAreas []entities.FishingArea, err error) {
	err = f.db.Joins("District").Select(selectedField).Find(&fishingAreas).Error
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
