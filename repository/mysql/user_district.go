package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserDistrictRepository interface {
	Create(userDistrict *entities.UserDistrict) error
	GetByUserID(userID int) (userDistrict entities.UserDistrict, err error)
}

type userDistrictRepository struct {
	db gorm.DB
}

func (u *userDistrictRepository) GetByUserID(userID int) (userDistrict entities.UserDistrict, err error) {
	err = u.db.Preload("District").Where("user_id = ?", userID).First(&userDistrict).Error
	if err != nil {
		return entities.UserDistrict{}, err
	}

	return userDistrict, nil
}

func (u *userDistrictRepository) Create(userDistrict *entities.UserDistrict) error {
	err := u.db.Create(&userDistrict).Error
	return err
}

func NewUserDistrictRepository(database gorm.DB) UserDistrictRepository {
	return &userDistrictRepository{db: database}
}
