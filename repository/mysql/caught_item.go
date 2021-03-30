package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type CaughtItemRepository interface {
	GetByID(id int) (entities.CaughtItem, error)
	Update(id int, data map[string]interface{}) error
	Search(query map[string]interface{}) ([]entities.CaughtItem, error)
	BulkUpdate(ids []int, data map[string]interface{}) error
	Index(query map[string]interface{}, date string) ([]entities.CaughtItem, error)
	Delete(id int) error
}

type caughtItemRepository struct {
	db gorm.DB
}

func (c *caughtItemRepository) Delete(id int) error {
	err := c.db.Delete(&entities.CaughtItem{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *caughtItemRepository) Index(query map[string]interface{}, date string) ([]entities.CaughtItem, error) {
	var result []entities.CaughtItem

	err := c.db.Table("caught_items").
		Joins("INNER JOIN caughts ON caughts.id = caught_items.caught_id").
		Where(query).
		Where("DATE(caughts.created_at) = DATE(?)", date).
		Preload("Caught.Fisher").
		Preload("Caught.FishingGear").
		Preload("Caught.FishingArea").
		Preload("FishType").
		Preload("CaughtStatus").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *caughtItemRepository) BulkUpdate(ids []int, data map[string]interface{}) error {
	err := c.db.Table("caught_items").Where("id IN ?", ids).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *caughtItemRepository) Search(query map[string]interface{}) ([]entities.CaughtItem, error) {
	var (
		result []entities.CaughtItem
		db     = c.db.Table("caught_items")
	)

	err := db.Joins("INNER JOIN caughts ON caughts.id = caught_items.caught_id").
		Where(query).
		Preload("Caught.Fisher").
		Preload("FishType").
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *caughtItemRepository) GetByID(id int) (entities.CaughtItem, error) {
	var caughtItem entities.CaughtItem

	err := c.db.Preload("Caught").Preload("FishType").Find(&caughtItem, id).Error
	if err != nil {
		return entities.CaughtItem{}, err
	}

	return caughtItem, nil
}

func (c *caughtItemRepository) Update(id int, data map[string]interface{}) error {
	err := c.db.Model(&entities.CaughtItem{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func NewCaughtItemRepository(database gorm.DB) CaughtItemRepository {
	return &caughtItemRepository{db: database}
}
