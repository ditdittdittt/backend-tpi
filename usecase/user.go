package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
)

type UserUsecase interface {
	Login(username string, password string) (token string, err error)
	Logout(id int) error
	GetUser(id int) (user entities.User, err error)
	Index() (users []entities.User, err error)
	Update(user *entities.User) error
	GetByID(id int) (entities.User, error)
}

type userUsecase struct {
	jwtService     services.JWTService
	userRepository mysql.UserRepository
}

func (u *userUsecase) Logout(id int) error {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] User repository error")
	}

	user.Token = ""
	err = u.userRepository.Update(&user)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] User repository error")
	}

	return nil
}

func (u *userUsecase) GetByID(id int) (entities.User, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return user, stacktrace.Propagate(err, "[GetByID] User repository error")
	}

	return user, nil
}

func (u *userUsecase) Update(user *entities.User) error {
	user.UpdatedAt = time.Now()

	err := u.userRepository.Update(user)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] User repository error")
	}

	return nil
}

func (u *userUsecase) Index() (users []entities.User, err error) {
	users, err = u.userRepository.Get()
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] User repository error")
	}

	return users, nil
}

func (u *userUsecase) GetUser(id int) (user entities.User, err error) {
	user, err = u.userRepository.GetByID(id)
	if err != nil {
		return entities.User{}, stacktrace.Propagate(err, "[GetByID] User repository error")
	}

	return user, err
}

func (u *userUsecase) Login(username string, password string) (token string, err error) {
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		return "", stacktrace.Propagate(err, "[GetByUsername] User repository error")
	}

	if user.Password != password {
		return "", stacktrace.NewError("Password didn't match")
	}

	token, err = u.jwtService.GenerateToken(&user)
	if err != nil {
		return "", stacktrace.Propagate(err, "[GenerateToken] Jwt Service error")
	}

	user.Token = token
	err = u.userRepository.Update(&user)
	if err != nil {
		return "", stacktrace.Propagate(err, "[Update] User repository error")
	}

	return token, nil
}

func NewUserUsecase(jwtService services.JWTService, userRepository mysql.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository, jwtService: jwtService}
}
