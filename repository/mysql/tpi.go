package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TpiRepository interface {
	Create(tpi *entities.Tpi) error
	GetByID(id int) (tpi entities.Tpi, err error)
}

type tpiRepository struct {
	db gorm.DB
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
