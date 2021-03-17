package usecase

import (
	"database/sql"
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
)

type UserUsecase interface {
	Login(username string, password string) (token string, err error)
	Logout(id int) error
	GetUser(id int) (entities.User, map[string]interface{}, error)
	Index() (users []entities.User, err error)
	Update(user *entities.User) error
	GetByID(id int) (entities.User, error)
}

type userUsecase struct {
	jwtService             services.JWTService
	userRepository         mysql.UserRepository
	userDistrictRepository mysql.UserDistrictRepository
	userTpiRepository      mysql.UserTpiRepository
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

	_, err := u.userRepository.GetByUsername(user.Username)
	if err != sql.ErrNoRows {
		return stacktrace.NewError("Username already used")
	}

	err = u.userRepository.Update(user)
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

func (u *userUsecase) GetUser(id int) (entities.User, map[string]interface{}, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return entities.User{}, nil, stacktrace.Propagate(err, "[GetByID] User repository error")
	}

	for _, permission := range user.Role.Permission {
		user.Permissions = append(user.Permissions, permission.Name)
	}

	locationData := map[string]interface{}{}

	switch user.RoleID {
	case 2:
		userDetail, err := u.userDistrictRepository.GetByUserID(user.ID)
		if err != nil {
			return entities.User{}, nil, stacktrace.Propagate(err, "[GetByUserID] User district repository error")
		}
		locationData["location_name"] = userDetail.District.Name
		locationData["location_id"] = userDetail.DistrictID
		return user, locationData, nil
	case 3, 4, 5:
		userDetail, err := u.userTpiRepository.GetByUserID(user.ID)
		if err != nil {
			return entities.User{}, nil, stacktrace.Propagate(err, "[GetByUserID] User district repository error")
		}
		locationData["location_name"] = userDetail.Tpi.Name
		locationData["location_id"] = userDetail.TpiID
		return user, locationData, nil
	default:
		return user, nil, nil
	}

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

func NewUserUsecase(jwtService services.JWTService, userRepository mysql.UserRepository, userDistrictRepository mysql.UserDistrictRepository, userTpiRepository mysql.UserTpiRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository, jwtService: jwtService, userDistrictRepository: userDistrictRepository, userTpiRepository: userTpiRepository}
}
