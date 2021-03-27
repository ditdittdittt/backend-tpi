package services

import (
	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/entities"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&entities.Auction{},
		&entities.Buyer{},
		&entities.BuyerTpi{},
		&entities.Caught{},
		&entities.CaughtItem{},
		&entities.CaughtStatus{},
		&entities.District{},
		&entities.FishType{},
		&entities.Fisher{},
		&entities.FisherTpi{},
		&entities.FishingArea{},
		&entities.FishingGear{},
		&entities.Permission{},
		&entities.Province{},
		&entities.Role{},
		&entities.Tpi{},
		&entities.Transaction{},
		&entities.TransactionItem{},
		&entities.User{},
		&entities.UserDistrict{},
		&entities.UserStatus{},
		&entities.UserSuperadmin{},
		&entities.UserTpi{},
	)

	if err != nil {
		panic(err)
	}
}
