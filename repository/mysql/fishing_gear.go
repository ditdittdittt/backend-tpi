package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishingGearRepository interface {
	Create(fishingGear *entities.FishingGear) error
	GetWithSelectedField(selectedField []string, queryMap map[string]interface{}) (fishingGears []entities.FishingGear, err error)
	GetByID(id int) (fishingGear entities.FishingGear, err error)
	Update(fishingGear *entities.FishingGear) error
	Delete(id int) error
}

type fishingGearRepository struct {
	db gorm.DB
}

func (f *fishingGearRepository) GetByID(id int) (fishingGear entities.FishingGear, err error) {
	err = f.db.First(&fishingGear, id).Error
	if err != nil {
		return entities.FishingGear{}, err
	}
	return fishingGear, err
}

func (f *fishingGearRepository) Update(fishingGear *entities.FishingGear) error {
	err := f.db.Model(&fishingGear).Updates(fishingGear).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *fishingGearRepository) Delete(id int) error {
	err := f.db.Delete(&entities.FishingGear{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (f *fishingGearRepository) GetWithSelectedField(selectedField []string, queryMap map[string]interface{}) (fishingGears []entities.FishingGear, err error) {
	err = f.db.Joins("District").Select(selectedField).Find(&fishingGears, queryMap).Error
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
