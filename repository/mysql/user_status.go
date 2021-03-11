package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserStatusRepository interface {
	Create(userStatus *entities.UserStatus) error
}

type userStatusRepository struct {
	db gorm.DB
}

func (u *userStatusRepository) Create(userStatus *entities.UserStatus) error {
	err := u.db.Create(&userStatus).Error
	return err
}

func NewUserStatusRepository(database gorm.DB) UserStatusRepository {
	return &userStatusRepository{db: database}
}
