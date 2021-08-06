package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	configuration "github.com/ditdittdittt/backend-tpi/config"
	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/delivery/http"
	"github.com/ditdittdittt/backend-tpi/repository/client"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

func main() {
	configuration.Init()
	database.Init()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			services.Migrate()
			break
		case "seed":
			services.Seed()
			break
		case "test":
			services.Test()
			break
		}
		return
	}

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	r.Use(cors.New(config))
	jwtService := services.NewJWTAuthService()

	// TPI
	tpiRepository := mysql.NewTpiRepository(*database.DB)
	tpiUsecase := usecase.NewTpiUsecase(tpiRepository)
	http.NewTpiHandler(r, tpiUsecase)

	userDistrictRepository := mysql.NewUserDistrictRepository(*database.DB)
	userTpiRepository := mysql.NewUserTpiRepository(*database.DB)
	userRepository := mysql.NewUserRepository(*database.DB)

	// User District
	userDistrictUsecase := usecase.NewUserDistrictUsecase(userDistrictRepository, userRepository)
	http.NewUserDistrictHandler(r, userDistrictUsecase)

	// User Tpi
	userTpiUsecase := usecase.NewUserTpiUsecase(userTpiRepository, userRepository)
	http.NewUserTpiHandler(r, userTpiUsecase)

	// User
	userUsecase := usecase.NewUserUsecase(jwtService, userRepository, userDistrictRepository, userTpiRepository, tpiRepository)
	http.NewUserHandler(r, userUsecase)

	// Fisher
	fisherRepository := mysql.NewFisherRepository(*database.DB)
	fisherTpiRepository := mysql.NewFisherTpiRepository(*database.DB)
	fisherUsecase := usecase.NewFisherUsecase(fisherRepository, fisherTpiRepository)
	http.NewFisherHandler(r, fisherUsecase)

	// Buyer
	buyerRepository := mysql.NewBuyerRepository(*database.DB)
	buyerTpiRepository := mysql.NewBuyerTpiRepository(*database.DB)
	buyerUsecase := usecase.NewBuyerUsecase(buyerRepository, buyerTpiRepository)
	http.NewBuyerHandler(r, buyerUsecase)

	// Fishing gear
	fishingGearRepository := mysql.NewFishingGearRepository(*database.DB)
	fishingGearUsecase := usecase.NewFishingGearUsecase(fishingGearRepository, tpiRepository)
	http.NewFishingGearHandler(r, fishingGearUsecase)

	// Fish type
	fishTypeRepository := mysql.NewFishTypeRepository(*database.DB)
	fishTypeUsecase := usecase.NewFishTypeUsecase(fishTypeRepository)
	http.NewFishTypeHandler(r, fishTypeUsecase)

	// Caught
	caughtRepository := mysql.NewCaughtRepository(*database.DB)
	caughtItemRepository := mysql.NewCaughtItemRepository(*database.DB)
	caughtUsecase := usecase.NewCaughtUsecase(caughtRepository, caughtItemRepository)
	http.NewCaughtHandler(r, caughtUsecase)

	// Auction
	auctionRepository := mysql.NewAuctionRepository(*database.DB)
	auctionUsecase := usecase.NewAuctionUsecase(auctionRepository, caughtRepository, caughtItemRepository)
	http.NewAuctionHandler(r, auctionUsecase)

	// Fishing area
	fishingAreaRepository := mysql.NewFishingAreaRepository(*database.DB)
	fishingAreaUsecase := usecase.NewFishingAreaUsecase(fishingAreaRepository, tpiRepository)
	http.NewFishingAreaHandler(r, fishingAreaUsecase)

	// District
	districtRepository := mysql.NewDistrictRepository(*database.DB)
	districtUsecase := usecase.NewDistrictUsecase(districtRepository)
	http.NewDistrictHandler(r, districtUsecase)

	// Province
	provinceClientRepository := client.NewProvinceRepository()
	provinceUsecase := usecase.NewProvinceUsecase(provinceClientRepository)
	http.NewProvinceHandler(r, provinceUsecase)

	// Transaction
	transactionItemRepository := mysql.NewTransactionItemRepository(*database.DB)
	transactionRepository := mysql.NewTransactionRepository(*database.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, auctionRepository, caughtRepository, transactionItemRepository, caughtItemRepository, tpiRepository)
	http.NewTransactionHandler(r, transactionUsecase)

	// Report
	reportUsecase := usecase.NewReportUsecase(caughtRepository, auctionRepository, transactionRepository, fishTypeRepository, transactionItemRepository, tpiRepository, districtRepository)
	http.NewReportHandler(r, reportUsecase)

	// Dashboard
	dashboardUsecase := usecase.NewDashboardUsecase(caughtRepository, auctionRepository, transactionRepository, tpiRepository)
	http.NewDashboardHandler(r, dashboardUsecase)

	r.Run(":9090")
}
