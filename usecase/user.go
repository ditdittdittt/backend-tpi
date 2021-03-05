package usecase

import (
	"fmt"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
	"github.com/ditdittdittt/backend-tpi/services"
)

type UserUsecase interface {
	Login(username string, password string) (token string, err error)
	GetUser(id int) (user entities.User, err error)
	Index() (users []entities.User, err error)
}

type userUsecase struct {
	jwtService     services.JWTService
	userRepository repository.UserRepository
}

func (u *userUsecase) Index() (users []entities.User, err error) {
	selectedField := []string{"username", "name", "nik", "address", "role_id", "status", "created_at", "updated_at"}

	users, err = u.userRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fisher repository error")
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
	fmt.Println(user.Role.Permission)

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

func NewUserUsecase(jwtService services.JWTService, userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository, jwtService: jwtService}
}
