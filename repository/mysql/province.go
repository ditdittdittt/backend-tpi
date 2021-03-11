package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type ProvinceRepository interface {
	BulkInsert(provinces []entities.Province) error
}

type provinceRepository struct {
	db gorm.DB
}

func (p *provinceRepository) BulkInsert(provinces []entities.Province) error {
	err := p.db.Create(&provinces).Error
	if err != nil {
		return err
	}

	return nil
}

func NewProvinceRepository(database gorm.DB) ProvinceRepository {
	return &provinceRepository{db: database}
}
