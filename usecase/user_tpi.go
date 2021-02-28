package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type UserTpiUsecase interface {
	CreateTpiAccount(userTpi *entities.UserTpi) error
}

type userTpiUsecase struct {
	userTpiUsecase repository.UserTpiRepository
}

func (u *userTpiUsecase) CreateTpiAccount(userTpi *entities.UserTpi) error {

	userTpi.User.UserStatusID = 1
	userTpi.User.Password = userTpi.User.Username
	userTpi.User.CreatedAt = time.Now()
	userTpi.User.UpdatedAt = time.Now()

	err := u.userTpiUsecase.Create(userTpi)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] User district repository error")
	}

	return nil
}

func NewUserTpiUsecase(userTpiRepository repository.UserTpiRepository) UserTpiUsecase {
	return &userTpiUsecase{userTpiUsecase: userTpiRepository}
}