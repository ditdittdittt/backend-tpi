package usecase

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type UserTpiUsecase interface {
	CreateTpiAccount(userTpi *entities.UserTpi) error
}

type userTpiUsecase struct {
	userTpiRepository mysql.UserTpiRepository
	userRepository    mysql.UserRepository
}

func (u *userTpiUsecase) CreateTpiAccount(userTpi *entities.UserTpi) error {

	existingUser, err := u.userRepository.GetByUsername(userTpi.User.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existingUser.Username == userTpi.User.Username {
		return errors.New("Username already exist")
	}

	userTpi.User.UserStatusID = 1
	userTpi.User.Password = helper.HashAndSaltPassword([]byte(userTpi.User.Username))
	userTpi.User.CreatedAt = time.Now()
	userTpi.User.UpdatedAt = time.Now()

	err = u.userTpiRepository.Create(userTpi)
	if err != nil {
		return err
	}

	return nil
}

func NewUserTpiUsecase(userTpiRepository mysql.UserTpiRepository, userRepository mysql.UserRepository) UserTpiUsecase {
	return &userTpiUsecase{userTpiRepository: userTpiRepository, userRepository: userRepository}
}
