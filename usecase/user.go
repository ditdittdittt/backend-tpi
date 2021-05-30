package usecase

import (
	"errors"
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
)

type UserUsecase interface {
	Login(username string, password string) (token string, err error)
	Logout(id int) error
	GetUser(id int) (entities.User, map[string]interface{}, error)
	Index(userID int, tpiID int, districtID int) (users []entities.User, err error)
	Update(user *entities.User) error
	GetByID(id int) (entities.User, error)
	ChangePassword(id int, oldPassword string, newPassword string) error
	ResetPassword(id int) error
}

type userUsecase struct {
	jwtService             services.JWTService
	userRepository         mysql.UserRepository
	userDistrictRepository mysql.UserDistrictRepository
	userTpiRepository      mysql.UserTpiRepository
	tpiRepository          mysql.TpiRepository
}

func (u *userUsecase) ResetPassword(id int) error {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return err
	}

	user.Password = helper.HashAndSaltPassword([]byte(user.Username))

	err = u.userRepository.Update(&user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) ChangePassword(id int, oldPassword string, newPassword string) error {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return err
	}

	if !helper.ComparePassword(user.Password, []byte(oldPassword)) {
		return errors.New("Password didn't match")
	}

	user.Password = helper.HashAndSaltPassword([]byte(newPassword))
	err = u.userRepository.Update(&user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Logout(id int) error {
	_, err := u.userRepository.GetByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetByID(id int) (entities.User, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUsecase) Update(user *entities.User) error {
	// insert log
	err := helper.InsertLog(user.ID, constant.User)
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Index(userID int, tpiID int, districtID int) (users []entities.User, err error) {
	if tpiID != 0 {
		usersTpi, err := u.userTpiRepository.GetByTpiIDs([]int{tpiID})
		if err != nil {
			return nil, err
		}
		for _, userTpi := range usersTpi {
			if userTpi.User.ID == userID {
				continue
			}
			users = append(users, *userTpi.User)
		}
	}

	if districtID != 0 {
		queryMap := map[string]interface{}{
			"district_id": districtID,
		}
		tpis, err := u.tpiRepository.Get(queryMap)
		if err != nil {
			return nil, err
		}

		usersTpi, err := u.userTpiRepository.GetByTpiIDs(tpiToID(tpis))
		if err != nil {
			return nil, err
		}
		for _, userTpi := range usersTpi {
			userTpi.User.LocationID = userTpi.TpiID
			userTpi.User.LocationName = userTpi.Tpi.Name
			if userTpi.User.ID == userID {
				continue
			}
			users = append(users, *userTpi.User)
		}
	}

	if tpiID == 0 && districtID == 0 {
		usersGet, err := u.userRepository.Get()
		if err != nil {
			return nil, err
		}

		for _, user := range usersGet {
			switch user.RoleID {
			case 2:
				locationDetail, err := u.userDistrictRepository.GetByUserID(user.ID)
				if err != nil {
					return nil, stacktrace.Propagate(err, "[GetByUserID] User district repository error")
				}
				user.LocationID = locationDetail.District.ID
				user.LocationName = locationDetail.District.Name
			case 3, 4, 5:
				locationDetail, err := u.userTpiRepository.GetByUserID(user.ID)
				if err != nil {
					return nil, stacktrace.Propagate(err, "[GetByUserID] User tpi repository error")
				}
				user.LocationID = locationDetail.Tpi.ID
				user.LocationName = locationDetail.Tpi.Name
			}

			users = append(users, user)
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
		return "", errors.New("Username not found")
	}

	if user.UserStatusID == 2 {
		return "", errors.New("account inactive")
	}

	if !helper.ComparePassword(user.Password, []byte(password)) {
		return "", errors.New("Password didn't match")
	}

	token, err = u.jwtService.GenerateToken(&user)
	if err != nil {
		return "", err
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
