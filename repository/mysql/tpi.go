package mysql

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TpiRepository interface {
	Create(tpi *entities.Tpi) error
	GetByID(id int) (tpi entities.Tpi, err error)
	Update(tpi *entities.Tpi) error
	Delete(id int) error
	Get(query map[string]interface{}) (tpis []entities.Tpi, err error)

	GetLatestCode(districtID int) (string, error)
}

type tpiRepository struct {
	db gorm.DB
}

func (t *tpiRepository) GetLatestCode(districtID int) (string, error) {
	var result string

	query := `SELECT code FROM tpis WHERE district_id = ` + strconv.Itoa(districtID) + ` ORDER BY code DESC LIMIT 1`

	err := t.db.Raw(query).Scan(&result).Error
	if err != nil {
		return "", err
	}

	return result, nil
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
	err = t.db.Preload("District").Find(&tpi, id).Error
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
