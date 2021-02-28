package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type DistrictRepository interface {
	Create(district *entities.District) error
}

type districtRepository struct {
	db gorm.DB
}

func (d *districtRepository) Create(district *entities.District) error {
	err := d.db.Create(&district).Error
	return err
}

func NewDistrictRepository(database gorm.DB) DistrictRepository {
	return &districtRepository{db: database}
}
