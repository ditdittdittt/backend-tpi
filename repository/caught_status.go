package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type CaughtStatusRepository interface {
	Create(caughtStatus *entities.CaughtStatus) error
}

type caughtStatusRepository struct {
	db gorm.DB
}

func (c *caughtStatusRepository) Create(caughtStatus *entities.CaughtStatus) error {
	err := c.db.Create(&caughtStatus).Error
	return err
}

func NewCaughtStatusRepository(database gorm.DB) CaughtStatusRepository {
	return &caughtStatusRepository{db: database}
}