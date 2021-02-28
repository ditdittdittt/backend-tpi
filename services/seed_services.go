package services

import (
	"time"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

func Seed() {
	//caughtStatusRepository := repository.NewCaughtStatusRepository(*database.DB)
	//caughtStatus1 := &entities.CaughtStatus{
	//	ID:     1,
	//	Status: "Belum terjual",
	//}
	//caughtStatus2 := &entities.CaughtStatus{
	//	ID:     2,
	//	Status: "Sudah terjual",
	//}
	//

	//
	//fishTypeRepository := repository.NewFishTypeRepository(*database.DB)
	//fishType1 := &entities.FishType{
	//	ID:   1,
	//	Name: "Net",
	//	Code: "FT01",
	//}

	districtRepository := repository.NewDistrictRepository(*database.DB)
	district1 := &entities.District{
		ID:         1,
		ProvinceID: 1,
		Name:       "Bekasi",
	}
	districtRepository.Create(district1)

	tpiRepository := repository.NewTpiRepository(*database.DB)
	tpi1 := &entities.Tpi{
		ID:         1,
		DistrictID: 1,
		District:   entities.District{},
		Name:       "TPI Bekasi",
		Location:   "Bekasi",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Code:       "B01",
	}
	tpiRepository.Create(tpi1)

	// User status
	userStatusRepository := repository.NewUserStatusRepository(*database.DB)
	userStatusMap := map[int]string{
		1: "Active",
		2: "Inactive",
	}
	for index, key := range userStatusMap {
		userStatus := &entities.UserStatus{
			ID: index,
			Status: key,
		}
		userStatusRepository.Create(userStatus)
	}

	// Role
	roleRepository := repository.NewRoleRepository(*database.DB)
	roleMap := map[int]string{
		1: "superadmin",
		2: "district-admin",
		3: "tpi-admin",
		4: "tpi-officer",
		5: "tpi-cashier",
	}
	for index, key := range roleMap {
		role := &entities.Role{
			ID:         index,
			Name:       key,
		}
		roleRepository.Create(role)
	}

	// Permission
	permissionRepository := repository.NewPermissionRepository(*database.DB)
	permissionMap := map[int]string{
		1: constant.CreateDistrictAdmin,
		2: constant.CreateTpiAdmin,
		3: constant.CreateTpiOfficer,
		4: constant.CreateTpiCashier,
		5: constant.GetUser,
	}
	for index, key := range permissionMap {
		permission := &entities.Permission{
			ID:   index,
			Name: key,
		}
		permissionRepository.Create(permission)
	}

	// User
	userSuperadminRepository := repository.NewUserSuperadminRepository(*database.DB)
	userSuperadmin := &entities.UserSuperadmin{
		ID:     1,
		UserID: 1,
		User:   entities.User{
			ID:           1,
			RoleID:       1,
			Role:         entities.Role{
				ID:         1,
				Name:       "superadmin",
				Permission: []entities.Permission{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
					{
						ID: 3,
					},
					{
						ID: 4,
					},
					{
						ID: 5,
					},
				},
			},
			UserStatusID: 1,
			Nik:          "1234567890",
			Name:         "superadmin",
			Address:      "Bekasi",
			Username:     "superadmin",
			Password:     "superadmin",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Token:        "",
		},
	}
	userSuperadminRepository.Create(userSuperadmin)

}