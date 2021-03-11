package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/delivery/http"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

func main() {
	database.Init()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			services.Migrate()
			break
		case "seed":
			services.Seed()
			break
		}
		return
	}

	r := gin.Default()
	r.Use(cors.Default())
	// User
	jwtService := services.NewJWTAuthService()
	userRepository := mysql.NewUserRepository(*database.DB)
	userUsecase := usecase.NewUserUsecase(jwtService, userRepository)
	http.NewUserHandler(r, userUsecase)

	// User District
	userDistrictRepository := mysql.NewUserDistrictRepository(*database.DB)
	userDistrictUsecase := usecase.NewUserDistrictUsecase(userDistrictRepository)
	http.NewUserDistrictHandler(r, userDistrictUsecase)

	// User Tpi
	userTpiRepository := mysql.NewUserTpiRepository(*database.DB)
	userTpiUsecase := usecase.NewUserTpiUsecase(userTpiRepository)
	http.NewUserTpiHandler(r, userTpiUsecase)

	// Fisher
	fisherRepository := mysql.NewFisherRepository(*database.DB)
	fisherUsecase := usecase.NewFisherUsecase(fisherRepository)
	http.NewFisherHandler(r, fisherUsecase)

	// Buyer
	buyerRepository := mysql.NewBuyerRepository(*database.DB)
	buyerUsecase := usecase.NewBuyerUsecase(buyerRepository)
	http.NewBuyerHandler(r, buyerUsecase)

	// Fishing gear
	fishingGearRepository := mysql.NewFishingGearRepository(*database.DB)
	fishingGearUsecase := usecase.NewFishingGearUsecase(fishingGearRepository)
	http.NewFishingGearHandler(r, fishingGearUsecase)

	// Fishing area
	fishingAreaRepository := mysql.NewFishingAreaRepository(*database.DB)
	fishingAreaUsecase := usecase.NewFishingAreaUsecase(fishingAreaRepository)
	http.NewFishingAreaHandler(r, fishingAreaUsecase)

	// Fish type
	fishTypeRepository := mysql.NewFishTypeRepository(*database.DB)
	fishTypeUsecase := usecase.NewFishTypeUsecase(fishTypeRepository)
	http.NewFishTypeHandler(r, fishTypeUsecase)

	// Caught
	caughtRepository := mysql.NewCaughtRepository(*database.DB)
	caughtUsecase := usecase.NewCaughtUsecase(caughtRepository)
	http.NewCaughtHandler(r, caughtUsecase)

	// Auction
	auctionRepository := mysql.NewAuctionRepository(*database.DB)
	auctionUsecase := usecase.NewAuctionUsecase(auctionRepository, caughtRepository)
	http.NewAuctionHandler(r, auctionUsecase)

	// Transaction
	transactionRepository := mysql.NewTransactionRepository(*database.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, auctionRepository, caughtRepository)
	http.NewTransactionHandler(r, transactionUsecase)

	r.Run(":9090")
}
