package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FishTypeRepository interface {
	Create(fishType *entities.FishType) error
	GetWithSelectedField(selectedField []string) (fishTypes []entities.FishType, err error)
	GetByID(id int) (fishType entities.FishType, err error)
	Update(fishType *entities.FishType) error
	Delete(id int) error
}

type fishTypeRepository struct {
	db gorm.DB
}

func (f *fishTypeRepository) GetByID(id int) (fishType entities.FishType, err error) {
	err = f.db.First(&fishType, id).Error
	if err != nil {
		return entities.FishType{}, err
	}
	return fishType, nil
}

func (f *fishTypeRepository) Update(fishType *entities.FishType) error {
	err := f.db.Model(&fishType).Updates(fishType).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *fishTypeRepository) Delete(id int) error {
	err := f.db.Delete(&entities.FishType{}, id).Error
	if err != nil {
		return err
	}

	return nil
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
