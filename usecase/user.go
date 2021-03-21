package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
)

type UserUsecase interface {
	Login(username string, password string) (token string, err error)
	Logout(id int) error
	GetUser(id int) (entities.User, map[string]interface{}, error)
	Index(tpiID int, districtID int) (users []entities.User, err error)
	Update(user *entities.User) error
	GetByID(id int) (entities.User, error)
	ChangePassword(id int, oldPassword string, newPassword string) error
}

type userUsecase struct {
	jwtService             services.JWTService
	userRepository         mysql.UserRepository
	userDistrictRepository mysql.UserDistrictRepository
	userTpiRepository      mysql.UserTpiRepository
	tpiRepository          mysql.TpiRepository
}

func (u *userUsecase) ChangePassword(id int, oldPassword string, newPassword string) error {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] User repository error")
	}

	if !helper.ComparePassword(user.Password, []byte(oldPassword)) {
		return stacktrace.NewError("Password didn't match")
	}

	user.Password = helper.HashAndSaltPassword([]byte(newPassword))
	err = u.userRepository.Update(&user)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] User repository error")
	}

	return nil
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

func (u *userUsecase) Index(tpiID int, districtID int) (users []entities.User, err error) {
	if tpiID != 0 {
		usersTpi, err := u.userTpiRepository.GetByTpiIDs([]int{tpiID})
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByTpiID] User tpi repository error")
		}
		for _, userTpi := range usersTpi {
			userTpi.User.Token = ""
			users = append(users, userTpi.User)
		}
	}

	if districtID != 0 {
		queryMap := map[string]interface{}{
			"district_id": districtID,
		}
		tpis, err := u.tpiRepository.Get(queryMap)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[Get] Tpi repository error")
		}

		usersTpi, err := u.userTpiRepository.GetByTpiIDs(tpiToID(tpis))
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByTpiID] User tpi repository error")
		}
		for _, userTpi := range usersTpi {
			userTpi.User.Token = ""
			users = append(users, userTpi.User)
		}
	}

	if tpiID == 0 && districtID == 0 {
		users, err = u.userRepository.Get()
		if err != nil {
			return nil, stacktrace.Propagate(err, "[Get] User repository error")
		}
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

	if !helper.ComparePassword(user.Password, []byte(password)) {
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

func NewUserUsecase(
	jwtService services.JWTService,
	userRepository mysql.UserRepository,
	userDistrictRepository mysql.UserDistrictRepository,
	userTpiRepository mysql.UserTpiRepository,
	tpiRepository mysql.TpiRepository,
) UserUsecase {
	return &userUsecase{
		userRepository:         userRepository,
		jwtService:             jwtService,
		userDistrictRepository: userDistrictRepository,
		userTpiRepository:      userTpiRepository,
		tpiRepository:          tpiRepository,
	}
}

func tpiToID(tpis []entities.Tpi) []int {
	result := make([]int, 0)
	for _, tpi := range tpis {
		result = append(result, tpi.ID)
	}
	return result
}
