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

	r.Run(":9090")
}