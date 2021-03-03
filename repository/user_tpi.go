package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserTpiRepository interface {
	Create(userTpi *entities.UserTpi) error
	GetByUserID(userID int) (userTpi entities.UserTpi, err error)
}

type userTpiRepository struct {
	db gorm.DB
}

func (u *userTpiRepository) GetByUserID(userID int) (userTpi entities.UserTpi, err error) {
	err = u.db.Where("user_id = ?", userID).First(&userTpi).Error
	if err != nil {
		return entities.UserTpi{}, err
	}

	return userTpi, nil
}

func (u *userTpiRepository) Create(userTpi *entities.UserTpi) error {
	err := u.db.Create(&userTpi).Error
	return err
}

func NewUserTpiRepository(database gorm.DB) UserTpiRepository {
	return &userTpiRepository{db: database}
}
