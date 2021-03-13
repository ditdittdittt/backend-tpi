package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TpiRepository interface {
	Create(tpi *entities.Tpi) error
	GetByID(id int) (tpi entities.Tpi, err error)
	Update(tpi *entities.Tpi) error
	Delete(id int) error
	Get(query map[string]interface{}) (tpis []entities.Tpi, err error)
}

type tpiRepository struct {
	db gorm.DB
}

func (t *tpiRepository) Update(tpi *entities.Tpi) error {
	err := t.db.Model(&tpi).Updates(tpi).Error
	if err != nil {
		return err
	}
	return nil

}

func (t *tpiRepository) Delete(id int) error {
	err := t.db.Delete(&entities.Tpi{}, id).Error
	if err != nil {
		return err
	}

	return nil

}

func (t *tpiRepository) Get(query map[string]interface{}) (tpis []entities.Tpi, err error) {
	err = t.db.Find(&tpis, query).Error
	if err != nil {
		return nil, err
	}
	return tpis, err

}

func (t *tpiRepository) GetByID(id int) (tpi entities.Tpi, err error) {
	err = t.db.Find(&tpi, id).Error
	if err != nil {
		return entities.Tpi{}, err
	}

	return tpi, nil
}

func (t *tpiRepository) Create(tpi *entities.Tpi) error {
	err := t.db.Create(&tpi).Error
	return err
}

func NewTpiRepository(database gorm.DB) TpiRepository {
	return &tpiRepository{db: database}
}
