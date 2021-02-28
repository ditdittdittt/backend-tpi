package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserTpiRepository interface {
	Create(userTpi *entities.UserTpi) error
}

type userTpiRepository struct {
	db gorm.DB
}

func (u *userTpiRepository) Create(userTpi *entities.UserTpi) error {
	err := u.db.Create(&userTpi).Error
	return err
}

func NewUserTpiRepository(database gorm.DB) UserTpiRepository {
	return &userTpiRepository{db: database}
}