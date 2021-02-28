package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type UserDistrictUsecase interface {
	CreateDistrictAccount(userDistrict *entities.UserDistrict) error
}

type userDistrictUsecase struct {
	userDistrictRepository repository.UserDistrictRepository
}

func (u *userDistrictUsecase) CreateDistrictAccount(userDistrict *entities.UserDistrict) error {

	userDistrict.User.UserStatusID = 1
	userDistrict.User.Password = userDistrict.User.Username
	userDistrict.User.CreatedAt = time.Now()
	userDistrict.User.UpdatedAt = time.Now()

	err := u.userDistrictRepository.Create(userDistrict)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] User district repository error")
	}

	return nil
}

func NewUserDistrictUsecase(userDistrictRepository repository.UserDistrictRepository) UserDistrictUsecase {
	return &userDistrictUsecase{userDistrictRepository: userDistrictRepository}
}
