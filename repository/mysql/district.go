package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type DistrictRepository interface {
	Create(district *entities.District) error
	BulkInsert(districts []entities.District) error
	Get(query map[string]interface{}) (districts []entities.District, err error)
}

type districtRepository struct {
	db gorm.DB
}

func (d *districtRepository) Get(query map[string]interface{}) (districts []entities.District, err error) {
	err = d.db.Find(&districts, query).Error
	if err != nil {
		return nil, err
	}

	return districts, nil
}

func (d *districtRepository) BulkInsert(districts []entities.District) error {
	err := d.db.Create(&districts).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *districtRepository) Create(district *entities.District) error {
	err := d.db.Create(&district).Error
	return err
}

func NewDistrictRepository(database gorm.DB) DistrictRepository {
	return &districtRepository{db: database}
}
