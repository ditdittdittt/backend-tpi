package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/delivery/http"
	"github.com/ditdittdittt/backend-tpi/repository/client"
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
	jwtService := services.NewJWTAuthService()

	// TPI
	tpiRepository := mysql.NewTpiRepository(*database.DB)
	tpiUsecase := usecase.NewTpiUsecase(tpiRepository)
	http.NewTpiHandler(r, tpiUsecase)

	// User District
	userDistrictRepository := mysql.NewUserDistrictRepository(*database.DB)
	userDistrictUsecase := usecase.NewUserDistrictUsecase(userDistrictRepository)
	http.NewUserDistrictHandler(r, userDistrictUsecase)

	// User Tpi
	userTpiRepository := mysql.NewUserTpiRepository(*database.DB)
	userTpiUsecase := usecase.NewUserTpiUsecase(userTpiRepository)
	http.NewUserTpiHandler(r, userTpiUsecase)

	// User
	userRepository := mysql.NewUserRepository(*database.DB)
	userUsecase := usecase.NewUserUsecase(jwtService, userRepository, userDistrictRepository, userTpiRepository, tpiRepository)
	http.NewUserHandler(r, userUsecase)

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
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, auctionRepository, caughtRepository, transactionItemRepository)
	http.NewTransactionHandler(r, transactionUsecase)

	// Report
	reportUsecase := usecase.NewReportUsecase(caughtRepository, auctionRepository, transactionRepository, fishTypeRepository, transactionItemRepository, tpiRepository)
	http.NewReportHandler(r, reportUsecase)

	// Dashboard
	dashboardUsecase := usecase.NewDashboardUsecase(caughtRepository, auctionRepository, transactionRepository)
	http.NewDashboardHandler(r, dashboardUsecase)

	r.Run(":9090")
}
