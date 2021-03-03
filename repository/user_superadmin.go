package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserSuperadminRepository interface {
	Create(userSuperadmin *entities.UserSuperadmin) error
	GetByUserID(userID int) (userSuperadmin entities.UserSuperadmin, err error)
}

type userSuperadminRepository struct {
	db gorm.DB
}

func (u *userSuperadminRepository) GetByUserID(userID int) (userSuperadmin entities.UserSuperadmin, err error) {
	err = u.db.Where("user_id = ?", userID).First(&userSuperadmin).Error
	if err != nil {
		return entities.UserSuperadmin{}, err
	}

	return userSuperadmin, nil
}

func (u *userSuperadminRepository) Create(userSuperadmin *entities.UserSuperadmin) error {
	err := u.db.Create(&userSuperadmin).Error
	return err
}

func NewUserSuperadminRepository(database gorm.DB) UserSuperadminRepository {
	return &userSuperadminRepository{db: database}
}
