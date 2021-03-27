package http

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateTpiAdminRequest struct {
	TpiID    int    `json:"tpi_id" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateTpiOfficerRequest struct {
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateTpiCashierRequest struct {
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateDistrictAdminRequest struct {
	DistrictID int    `json:"district_id" binding:"required"`
	Nik        string `json:"nik" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Username   string `json:"username" binding:"required"`
}

type CreateTpiRequest struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Pic         string `json:"pic"`
}

type CreateFisherRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	NickName    string `json:"nick_name"`
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
	SouthLatitudeDegree string `json:"south_latitude_degree"`
	SouthLatitudeMinute string `json:"south_latitude_minute"`
	SouthLatitudeSecond string `json:"south_latitude_second"`
	EastLongitudeDegree string `json:"east_longitude_degree"`
	EastLongitudeMinute string `json:"east_longitude_minute"`
	EastLongitudeSecond string `json:"east_longitude_second"`
	Name                string `json:"name"`
}

type CreateCaughtRequest struct {
	FisherID      int `json:"fisher_id"`
	TripDay       int `json:"trip_day"`
	FishingGearID int `json:"fishing_gear_id"`
	FishingAreaID int `json:"fishing_area_id"`
	CaughtItems   []struct {
		FishTypeID int     `json:"fish_type_id"`
		Weight     float64 `json:"weight"`
		WeightUnit string  `json:"weight_unit"`
	} `json:"caught_items"`
}

type CreateAuctionRequest struct {
	CaughtItemID int     `json:"caught_item_id"`
	Price        float64 `json:"price"`
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

type UpdateBuyerRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type UpdateFishingGearRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type UpdateFishingAreaRequest struct {
	DistrictID          int    `json:"district_id"`
	SouthLatitudeDegree string `json:"south_latitude_degree"`
	SouthLatitudeMinute string `json:"south_latitude_minute"`
	SouthLatitudeSecond string `json:"south_latitude_second"`
	EastLongitudeDegree string `json:"east_longitude_degree"`
	EastLongitudeMinute string `json:"east_longitude_minute"`
	EastLongitudeSecond string `json:"east_longitude_second"`
	Name                string `json:"name"`
}

type UpdateFishTypeRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type UpdateAuctionRequest struct {
	Price float64 `json:"price"`
}

type UpdateTpiRequest struct {
	DistrictID int    `json:"district_id"`
	Name       string `json:"name"`
}

type UpdateCaughtRequest struct {
	FisherID      int     `json:"fisher_id"`
	TripDay       int     `json:"trip_day"`
	FishingGearID int     `json:"fishing_gear_id"`
	FishingAreaID int     `json:"fishing_area_id"`
	FishTypeID    int     `json:"fish_type_id"`
	Weight        float64 `json:"weight"`
	WeightUnit    string  `json:"weight_unit"`
}

type UpdateTransactionRequest struct {
	BuyerID          int    `json:"buyer_id"`
	DistributionArea string `json:"distribution_area"`
}

type UpdateUserRequest struct {
	UserRoleID   int    `json:"user_role_id"`
	UserStatusID int    `json:"user_status_id"`
	Nik          string `json:"nik"`
	Name         string `json:"name"`
	Address      string `json:"address"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
