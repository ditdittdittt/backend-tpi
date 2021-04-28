package usecase

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type UserDistrictUsecase interface {
	CreateDistrictAccount(userDistrict *entities.UserDistrict) error
}

type userDistrictUsecase struct {
	userDistrictRepository mysql.UserDistrictRepository
	userRepository         mysql.UserRepository
}

func (u *userDistrictUsecase) CreateDistrictAccount(userDistrict *entities.UserDistrict) error {

	existingUser, err := u.userRepository.GetByUsername(userDistrict.User.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existingUser.Username == userDistrict.User.Username {
		return errors.New("Username already exist")
	}

	userDistrict.User.UserStatusID = 1
	userDistrict.User.Password = helper.HashAndSaltPassword([]byte(userDistrict.User.Username))
	userDistrict.User.CreatedAt = time.Now()
	userDistrict.User.UpdatedAt = time.Now()

	err = u.userDistrictRepository.Create(userDistrict)
	if err != nil {
		return err
	}

	return nil
}

func NewUserDistrictUsecase(userDistrictRepository mysql.UserDistrictRepository, userRepository mysql.UserRepository) UserDistrictUsecase {
	return &userDistrictUsecase{userDistrictRepository: userDistrictRepository, userRepository: userRepository}
}
