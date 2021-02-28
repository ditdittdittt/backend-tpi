package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserDistrictRepository interface {
	Create(userDistrict *entities.UserDistrict) error
}

type userDistrictRepository struct {
	db	gorm.DB
}

func (u *userDistrictRepository) Create(userDistrict *entities.UserDistrict) error {
	err := u.db.Create(&userDistrict).Error
	return err
}

func NewUserDistrictRepository(database gorm.DB) UserDistrictRepository {
	return &userDistrictRepository{db: database}
}