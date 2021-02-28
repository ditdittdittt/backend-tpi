package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type CaughtRepository interface {
	Create(caught *entities.Caught) error
}

type caughtRepository struct {
	db gorm.DB
}

func (c *caughtRepository) Create(caught *entities.Caught) error {
	err := c.db.Create(&caught).Error
	return err
}

func NewCaughtRepository(database gorm.DB) CaughtRepository {
	return &caughtRepository{db: database}
}
