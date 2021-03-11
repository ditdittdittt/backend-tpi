package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type CaughtRepository interface {
	GetByID(id int) (caught entities.Caught, err error)
	Create(caught *entities.Caught) error
	Update(caught *entities.Caught, data map[string]interface{}) error
	BulkUpdate(id []int, data map[string]interface{}) error
	Get() (caughts []entities.Caught, err error)
	Search(query map[string]interface{}) (caughts []entities.Caught, err error)
}

type caughtRepository struct {
	db gorm.DB
}

func (c *caughtRepository) Search(query map[string]interface{}) (caughts []entities.Caught, err error) {
	err = c.db.Preload("FishType").Preload("Fisher").Find(&caughts, query).Error
	if err != nil {
		return nil, err
	}
	return caughts, nil
}

func (c *caughtRepository) Get() (caughts []entities.Caught, err error) {
	err = c.db.Preload("Fisher").Preload("FishType").Preload("FishingGear").Preload("FishingArea").Preload("CaughtStatus").Find(&caughts).Error
	if err != nil {
		return nil, err
	}

	return caughts, nil
}

func (c *caughtRepository) BulkUpdate(id []int, data map[string]interface{}) error {
	err := c.db.Table("caughts").Where("id IN ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *caughtRepository) GetByID(id int) (caught entities.Caught, err error) {
	err = c.db.First(&caught, id).Error
	if err != nil {
		return entities.Caught{}, err
	}
	return caught, nil
}

func (c *caughtRepository) Update(caught *entities.Caught, data map[string]interface{}) error {
	err := c.db.Model(caught).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *caughtRepository) Create(caught *entities.Caught) error {
	err := c.db.Create(&caught).Error
	return err
}

func NewCaughtRepository(database gorm.DB) CaughtRepository {
	return &caughtRepository{db: database}
}
