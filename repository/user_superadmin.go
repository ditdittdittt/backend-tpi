package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserSuperadminRepository interface {
	Create(userSuperadmin *entities.UserSuperadmin) error
}

type userSuperadminRepository struct {
	db gorm.DB
}

func (u *userSuperadminRepository) Create(userSuperadmin *entities.UserSuperadmin) error {
	err := u.db.Create(&userSuperadmin).Error
	return err
}

func NewUserSuperadminRepository(database gorm.DB) UserSuperadminRepository {
	return &userSuperadminRepository{db: database}
}