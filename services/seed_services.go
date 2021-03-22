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
	seedCaughtStatus()
	seedFishType()
	seedFishingGear()
	seedProvinceAndDistrict()
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
		1: "Proses lelang",
		2: "Menunggu pembayaran",
		3: "Transaksi selesai",
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
		ID:   1,
		Name: "Net",
		Code: "FG01",
	}
	fishingGearRepository.Create(fishingGear1)
}

func seedUserStatus() {
	userStatusRepository := mysql.NewUserStatusRepository(*database.DB)
	userStatusMap := map[int]string{
		1: "Active",
		2: "Inactive",
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
		6:  constant.GetByIDUser,
		7:  constant.UpdateUser,
		8:  constant.CreateTpi,
		9:  constant.GetByIDTpi,
		10: constant.UpdateTpi,
		11: constant.DeleteTpi,
		12: constant.CreateCaught,
		13: constant.InquiryCaught,
		14: constant.GetByIDCaught,
		15: constant.UpdateCaught,
		16: constant.DeleteCaught,
		17: constant.CreateAuction,
		18: constant.InquiryAuction,
		19: constant.GetByIDAuction,
		20: constant.UpdateAuction,
		21: constant.DeleteAuction,
		22: constant.CreateTransaction,
		23: constant.GetByIDTransaction,
		24: constant.UpdateTransaction,
		25: constant.DeleteTransaction,
		26: constant.CreateFisher,
		27: constant.UpdateFisher,
		28: constant.GetByIDFisher,
		29: constant.DeleteFisher,
		30: constant.CreateBuyer,
		31: constant.UpdateBuyer,
		32: constant.GetByIDBuyer,
		33: constant.DeleteBuyer,
		34: constant.CreateFishingGear,
		35: constant.UpdateFishingGear,
		36: constant.GetByIDFishingGear,
		37: constant.DeleteFishingGear,
		38: constant.CreateFishingArea,
		39: constant.UpdateFishingArea,
		40: constant.GetByIDFishingArea,
		41: constant.DeleteFishingArea,
		42: constant.CreateFishType,
		43: constant.UpdateFishType,
		44: constant.GetByIDFishType,
		45: constant.DeleteFishType,
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
	permissionSuperadmin := []int{1, 2, 3, 4, 5, 6, 7, 42, 43, 44, 45}
	for _, permissionID := range permissionSuperadmin {
		permission := &entities.Permission{ID: permissionID}
		roleSuperadmin.Permission = append(roleSuperadmin.Permission, permission)
	}
	roleRepository.Create(roleSuperadmin)

	roleDistrictAdmin := &entities.Role{
		ID:   2,
		Name: "district-admin",
	}
	permissionDistrictAdmin := []int{2, 5, 6, 7, 8, 9, 10, 11, 14, 19, 23, 38, 39, 40, 41}
	for _, permissionID := range permissionDistrictAdmin {
		permission := &entities.Permission{ID: permissionID}
		roleDistrictAdmin.Permission = append(roleDistrictAdmin.Permission, permission)
	}
	roleRepository.Create(roleDistrictAdmin)

	roleTpiAdmin := &entities.Role{
		ID:   3,
		Name: "tpi-admin",
	}
	permissionTpiAdmin := []int{3, 4, 5, 6, 7, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 44}
	for _, permissionID := range permissionTpiAdmin {
		permission := &entities.Permission{ID: permissionID}
		roleTpiAdmin.Permission = append(roleTpiAdmin.Permission, permission)
	}
	roleRepository.Create(roleTpiAdmin)

	roleTpiOfficer := &entities.Role{
		ID:   4,
		Name: "tpi-officer",
	}
	permissionTpiOfficer := []int{5, 6, 7, 12, 13, 14, 17, 18, 19, 22, 23, 26, 27, 28, 30, 31, 32, 34, 36, 40, 44}
	for _, permissionID := range permissionTpiOfficer {
		permission := &entities.Permission{ID: permissionID}
		roleTpiOfficer.Permission = append(roleTpiOfficer.Permission, permission)
	}
	roleRepository.Create(roleTpiOfficer)

	roleTpiCashier := &entities.Role{
		ID:   5,
		Name: "tpi-cashier",
	}
	permissionTpiCashier := []int{5, 6, 7, 18, 22, 23, 24, 25}
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
		User: entities.User{
			ID:           1,
			RoleID:       1,
			UserStatusID: 1,
			Nik:          "1234567890",
			Name:         "superadmin",
			Address:      "Bekasi",
			Username:     "superadmin",
			Password:     helper.HashAndSaltPassword([]byte("superadmin")),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Token:        "",
		},
	}
	userSuperadminRepository.Create(userSuperadmin)
}
