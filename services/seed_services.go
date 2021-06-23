package services

import (
	"time"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/client"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

func Seed() {
	seedProvinceAndDistrict()
	seedCaughtStatus()
	seedFishType()
	seedFishingGear()
	seedUserStatus()
	seedRoleAndPermission()
	seedSuperadmin()

	//fishingAreaRepository := mysql.NewFishingAreaRepository(*database.DB)
	//fishingArea1 := &entities.FishingArea{
	//	ID:                  1,
	//	DistrictID:          32,
	//	Name:                "Bekasi",
	//	SouthLatitudeDegree: "145",
	//	SouthLatitudeMinute: "23",
	//	SouthLatitudeSecond: "12",
	//	EastLongitudeDegree: "123",
	//	EastLongitudeMinute: "23",
	//	EastLongitudeSecond: "12",
	//}
	//fishingAreaRepository.Create(fishingArea1)
	//
	//tpiRepository := mysql.NewTpiRepository(*database.DB)
	//tpi1 := &entities.Tpi{
	//	ID:         1,
	//	DistrictID: 1,
	//	Name:       "TPI Bekasi",
	//	CreatedAt:  time.Now(),
	//	UpdatedAt:  time.Now(),
	//	Code:       "B01",
	//}
	//tpiRepository.Create(tpi1)
	//
}

func seedProvinceAndDistrict() {
	// Province and District seed
	provinceClientRepository := client.NewProvinceRepository()
	provinceMysqlRepository := mysql.NewProvinceRepository(*database.DB)
	districtClientRepository := client.NewDistrictRepository()
	districtMysqlRepository := mysql.NewDistrictRepository(*database.DB)

	provinces, _ := provinceClientRepository.Get()
	provinceMysqlRepository.BulkInsert(provinces)

	for _, province := range provinces {
		districts, _ := districtClientRepository.GetByProvinceID(province.ID)
		districtMysqlRepository.BulkInsert(districts)
	}
}

func seedCaughtStatus() {
	caughtStatusRepository := mysql.NewCaughtStatusRepository(*database.DB)
	caughtStatusMap := map[int]string{
		1: "Belum Terjual",
		2: "Menunggu Pembayaran",
		3: "Selesai",
	}
	for index, key := range caughtStatusMap {
		caughtStatus := &entities.CaughtStatus{
			ID:     index,
			Status: key,
		}
		caughtStatusRepository.Create(caughtStatus)
	}
}

func seedFishType() {
	fishTypeRepository := mysql.NewFishTypeRepository(*database.DB)
	fishType1 := &entities.FishType{
		ID:   1,
		Name: "Tenggiri",
		Code: "FT01",
	}
	fishTypeRepository.Create(fishType1)
}

func seedFishingGear() {
	fishingGearRepository := mysql.NewFishingGearRepository(*database.DB)
	fishingGear1 := &entities.FishingGear{
		ID:         1,
		Name:       "Net",
		Code:       "FG01",
		DistrictID: 3212,
	}
	fishingGearRepository.Create(fishingGear1)
}

func seedUserStatus() {
	userStatusRepository := mysql.NewUserStatusRepository(*database.DB)
	userStatusMap := map[int]string{
		1: "Aktif",
		2: "Tidak Aktif",
	}
	for index, key := range userStatusMap {
		userStatus := &entities.UserStatus{
			ID:     index,
			Status: key,
		}
		userStatusRepository.Create(userStatus)
	}
}

func seedRoleAndPermission() {
	// Permission
	permissionRepository := mysql.NewPermissionRepository(*database.DB)
	permissionMap := []string{
		1:  constant.CreateDistrictAdmin,
		2:  constant.CreateTpiAdmin,
		3:  constant.CreateTpiOfficer,
		4:  constant.CreateTpiCashier,
		5:  constant.GetUser,
		6:  constant.ReadUser,
		7:  constant.UpdateUser,
		8:  constant.CreateTpi,
		9:  constant.ReadTpi,
		10: constant.UpdateTpi,
		11: constant.DeleteTpi,
		12: constant.CreateCaught,
		13: constant.ReadCaught,
		14: constant.UpdateCaught,
		15: constant.DeleteCaught,
		16: constant.CreateAuction,
		17: constant.ReadAuction,
		18: constant.UpdateAuction,
		19: constant.DeleteAuction,
		20: constant.CreateTransaction,
		21: constant.ReadTransaction,
		22: constant.UpdateTransaction,
		23: constant.DeleteTransaction,
		24: constant.CreateFisher,
		25: constant.UpdateFisher,
		26: constant.ReadFisher,
		27: constant.DeleteFisher,
		28: constant.CreateBuyer,
		29: constant.UpdateBuyer,
		30: constant.ReadBuyer,
		31: constant.DeleteBuyer,
		32: constant.CreateFishingGear,
		33: constant.UpdateFishingGear,
		34: constant.ReadFishingGear,
		35: constant.DeleteFishingGear,
		36: constant.CreateFishingArea,
		37: constant.UpdateFishingArea,
		38: constant.ReadFishingArea,
		39: constant.DeleteFishingArea,
		40: constant.CreateFishType,
		41: constant.UpdateFishType,
		42: constant.ReadFishType,
		43: constant.DeleteFishType,
	}
	for index, key := range permissionMap {
		permission := &entities.Permission{
			ID:   index,
			Name: key,
		}
		permissionRepository.Create(permission)
	}

	// Role
	roleRepository := mysql.NewRoleRepository(*database.DB)

	roleSuperadmin := &entities.Role{
		ID:   1,
		Name: "superadmin",
	}
	permissionSuperadmin := []int{1, 2, 3, 4, 5, 6, 7, 40, 41, 42, 43}
	for _, permissionID := range permissionSuperadmin {
		permission := &entities.Permission{ID: permissionID}
		roleSuperadmin.Permission = append(roleSuperadmin.Permission, permission)
	}
	roleRepository.Create(roleSuperadmin)

	roleDistrictAdmin := &entities.Role{
		ID:   2,
		Name: "district-admin",
	}
	permissionDistrictAdmin := []int{2, 5, 6, 7, 8, 9, 10, 11, 32, 33, 34, 35, 36, 37, 38, 39}
	for _, permissionID := range permissionDistrictAdmin {
		permission := &entities.Permission{ID: permissionID}
		roleDistrictAdmin.Permission = append(roleDistrictAdmin.Permission, permission)
	}
	roleRepository.Create(roleDistrictAdmin)

	roleTpiAdmin := &entities.Role{
		ID:   3,
		Name: "tpi-admin",
	}
	permissionTpiAdmin := []int{3, 4, 5, 6, 7, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 42}
	for _, permissionID := range permissionTpiAdmin {
		permission := &entities.Permission{ID: permissionID}
		roleTpiAdmin.Permission = append(roleTpiAdmin.Permission, permission)
	}
	roleRepository.Create(roleTpiAdmin)

	roleTpiOfficer := &entities.Role{
		ID:   4,
		Name: "tpi-officer",
	}
	permissionTpiOfficer := []int{5, 6, 7, 12, 13, 16, 17, 20, 21, 24, 26, 28, 30, 32, 34, 38, 40, 42}
	for _, permissionID := range permissionTpiOfficer {
		permission := &entities.Permission{ID: permissionID}
		roleTpiOfficer.Permission = append(roleTpiOfficer.Permission, permission)
	}
	roleRepository.Create(roleTpiOfficer)

	roleTpiCashier := &entities.Role{
		ID:   5,
		Name: "tpi-cashier",
	}
	permissionTpiCashier := []int{5, 6, 7, 17, 20, 21, 26, 28, 30, 42}
	for _, permissionID := range permissionTpiCashier {
		permission := &entities.Permission{ID: permissionID}
		roleTpiCashier.Permission = append(roleTpiCashier.Permission, permission)
	}
	roleRepository.Create(roleTpiCashier)
}

func seedSuperadmin() {
	// User
	userSuperadminRepository := mysql.NewUserSuperadminRepository(*database.DB)
	userSuperadmin := &entities.UserSuperadmin{
		ID:     1,
		UserID: 1,
		User: &entities.User{
			ID:           1,
			RoleID:       1,
			UserStatusID: 1,
			Nik:          "3216021204980014",
			Name:         "Yudit Yudiarto",
			Address:      "Bekasi",
			Username:     "superadmin",
			Password:     helper.HashAndSaltPassword([]byte("superadmin")),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	userSuperadminRepository.Create(userSuperadmin)
}
