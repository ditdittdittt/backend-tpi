package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/delivery/http"
	"github.com/ditdittdittt/backend-tpi/repository"
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
	// User
	jwtService := services.NewJWTAuthService()
	userRepository := repository.NewUserRepository(*database.DB)
	userUsecase := usecase.NewUserUsecase(jwtService, userRepository)
	http.NewUserHandler(r, userUsecase)

	// User District
	userDistrictRepository := repository.NewUserDistrictRepository(*database.DB)
	userDistrictUsecase := usecase.NewUserDistrictUsecase(userDistrictRepository)
	http.NewUserDistrictHandler(r, userDistrictUsecase)

	// User Tpi
	userTpiRepository := repository.NewUserTpiRepository(*database.DB)
	userTpiUsecase := usecase.NewUserTpiUsecase(userTpiRepository)
	http.NewUserTpiHandler(r, userTpiUsecase)

	// Fisher
	fisherRepository := repository.NewFisherRepository(*database.DB)
	fisherUsecase := usecase.NewFisherUsecase(fisherRepository)
	http.NewFisherHandler(r, fisherUsecase)

	// Buyer
	buyerRepository := repository.NewBuyerRepository(*database.DB)
	buyerUsecase := usecase.NewBuyerUsecase(buyerRepository)
	http.NewBuyerHandler(r, buyerUsecase)

	// Fishing gear
	fishingGearRepository := repository.NewFishingGearRepository(*database.DB)
	fishingGearUsecase := usecase.NewFishingGearUsecase(fishingGearRepository)
	http.NewFishingGearHandler(r, fishingGearUsecase)

	// Fishing area
	fishingAreaRepository := repository.NewFishingAreaRepository(*database.DB)
	fishingAreaUsecase := usecase.NewFishingAreaUsecase(fishingAreaRepository)
	http.NewFishingAreaHandler(r, fishingAreaUsecase)

	// Fish type
	fishTypeRepository := repository.NewFishTypeRepository(*database.DB)
	fishTypeUsecase := usecase.NewFishTypeUsecase(fishTypeRepository)
	http.NewFishTypeHandler(r, fishTypeUsecase)

	// Caught
	caughtRepository := repository.NewCaughtRepository(*database.DB)
	caughtUsecase := usecase.NewCaughtUsecase(caughtRepository)
	http.NewCaughtHandler(r, caughtUsecase)

	// Auction
	auctionRepository := repository.NewAuctionRepository(*database.DB)
	auctionUsecase := usecase.NewAuctionUsecase(auctionRepository, caughtRepository)
	http.NewAuctionHandler(r, auctionUsecase)

	// Transaction
	transactionRepository := repository.NewTransactionRepository(*database.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, auctionRepository, caughtRepository)
	http.NewTransactionHandler(r, transactionUsecase)

	r.Run(":9090")
}
