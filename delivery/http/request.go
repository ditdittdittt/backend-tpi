package http

import "github.com/ditdittdittt/backend-tpi/entities"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserRequest struct {
}

type CreateTpiAdminRequest struct {
	RoleID   int    `json:"role_id" binding:"required,eq=3"`
	TpiID    int    `json:"tpi_id" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateTpiOfficerRequest struct {
	RoleID   int    `json:"role_id" binding:"required,eq=4"`
	TpiID    int    `json:"tpi_id" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateTpiCashierRequest struct {
	RoleID   int    `json:"role_id" binding:"required,eq=5"`
	TpiID    int    `json:"tpi_id" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateDistrictAdminRequest struct {
	RoleID     int    `json:"role_id" binding:"required,eq=2"`
	DistrictID int    `json:"district_id" binding:"required"`
	Nik        string `json:"nik" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Username   string `json:"username" binding:"required"`
}

type CreateTpiRequest struct {
	DistrictID int    `json:"district_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Code       string `json:"code" binding:"required"`
}

type CreateFisherRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	ShipType    string `json:"ship_type"`
	AbkTotal    int    `json:"abk_total"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type CreateBuyerRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type CreateFishTypeRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateFishingGearRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateFishingAreaRequest struct {
	DistrictID          int    `json:"district_id"`
	SouthLatitudeDegree string `json:"south_latitude_degree"`
	SouthLatitudeMinute string `json:"south_latitude_minute"`
	SouthLatitudeSecond string `json:"south_latitude_second"`
	EastLongitudeDegree string `json:"east_longitude_degree"`
	EastLongitudeMinute string `json:"east_longitude_minute"`
	EastLongitudeSecond string `json:"east_longitude_second"`
	Name                string `json:"name"`
}

type CreateCaughtRequest struct {
	FisherID       int                   `json:"fisher_id"`
	TripDay        int                   `json:"trip_day"`
	FishingGearID  int                   `json:"fishing_gear_id"`
	FishingAreaID  int                   `json:"fishing_area_id"`
	CaughtFishData []entities.CaughtData `json:"caught_fish_data"`
}

type CreateAuctionRequest struct {
	CaughtID int     `json:"caught_id"`
	Price    float64 `json:"price"`
}

type CreateTransactionRequest struct {
	BuyerID          int     `json:"buyer_id"`
	DistributionArea string  `json:"distribution_area"`
	TotalPrice       float64 `json:"total_price"`
	AuctionsIDs      []int   `json:"auction_ids"`
}

type UpdateFisherRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	ShipType    string `json:"ship_type"`
	AbkTotal    int    `json:"abk_total"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}
