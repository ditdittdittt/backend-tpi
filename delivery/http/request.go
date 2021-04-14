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
	TpiID    int    `json:"tpi_id" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateTpiCashierRequest struct {
	TpiID    int    `json:"tpi_id" binding:"required"`
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
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pic         string `json:"pic" binding:"required"`
}

type CreateFisherRequest struct {
	Nik         string `json:"nik" binding:"required"`
	Name        string `json:"name" binding:"required"`
	NickName    string `json:"nick_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	ShipType    string `json:"ship_type" binding:"required"`
	AbkTotal    int    `json:"abk_total" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type CreateBuyerRequest struct {
	Nik         string `json:"nik" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type CreateFishTypeRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type CreateFishingGearRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type CreateFishingAreaRequest struct {
	SouthLatitudeDegree string `json:"south_latitude_degree" binding:"required"`
	SouthLatitudeMinute string `json:"south_latitude_minute" binding:"required"`
	SouthLatitudeSecond string `json:"south_latitude_second" binding:"required"`
	EastLongitudeDegree string `json:"east_longitude_degree" binding:"required"`
	EastLongitudeMinute string `json:"east_longitude_minute" binding:"required"`
	EastLongitudeSecond string `json:"east_longitude_second" binding:"required"`
	Name                string `json:"name" binding:"required"`
}

type CreateCaughtRequest struct {
	FisherID      int `json:"fisher_id" binding:"required"`
	TripDay       int `json:"trip_day" binding:"required"`
	FishingGearID int `json:"fishing_gear_id" binding:"required"`
	FishingAreaID int `json:"fishing_area_id" binding:"required"`
	CaughtItems   []struct {
		FishTypeID int     `json:"fish_type_id" binding:"required"`
		Weight     float64 `json:"weight" binding:"required"`
		WeightUnit string  `json:"weight_unit" binding:"required"`
	} `json:"caught_items" binding:"required"`
}

type CreateAuctionRequest struct {
	CaughtItemID int     `json:"caught_item_id" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

type CreateTransactionRequest struct {
	BuyerID          int     `json:"buyer_id" binding:"required"`
	DistributionArea string  `json:"distribution_area" binding:"required"`
	TotalPrice       float64 `json:"total_price" binding:"required"`
	AuctionsIDs      []int   `json:"auction_ids" binding:"required"`
}

type UpdateFisherRequest struct {
	Nik         string `json:"nik" binding:"required"`
	Name        string `json:"name" binding:"required"`
	NickName    string `json:"nick_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	ShipType    string `json:"ship_type" binding:"required"`
	AbkTotal    int    `json:"abk_total" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateBuyerRequest struct {
	Nik         string `json:"nik" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateFishingGearRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type UpdateFishingAreaRequest struct {
	DistrictID          int    `json:"district_id" binding:"required"`
	SouthLatitudeDegree string `json:"south_latitude_degree" binding:"required"`
	SouthLatitudeMinute string `json:"south_latitude_minute" binding:"required"`
	SouthLatitudeSecond string `json:"south_latitude_second" binding:"required"`
	EastLongitudeDegree string `json:"east_longitude_degree" binding:"required"`
	EastLongitudeMinute string `json:"east_longitude_minute" binding:"required"`
	EastLongitudeSecond string `json:"east_longitude_second" binding:"required"`
	Name                string `json:"name" binding:"required"`
}

type UpdateFishTypeRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type UpdateAuctionRequest struct {
	Price float64 `json:"price" binding:"required"`
}

type UpdateTpiRequest struct {
	DistrictID  int    `json:"district_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pic         string `json:"pic" binding:"required"`
}

type UpdateCaughtRequest struct {
	FisherID      int     `json:"fisher_id" binding:"required"`
	TripDay       int     `json:"trip_day" binding:"required"`
	FishingGearID int     `json:"fishing_gear_id" binding:"required"`
	FishingAreaID int     `json:"fishing_area_id" binding:"required"`
	FishTypeID    int     `json:"fish_type_id" v`
	Weight        float64 `json:"weight" binding:"required"`
	WeightUnit    string  `json:"weight_unit" binding:"required"`
	CaughtID      int     `json:"caught_id" binding:"required"`
}

type UpdateTransactionRequest struct {
	BuyerID          int    `json:"buyer_id" binding:"required"`
	DistributionArea string `json:"distribution_area" binding:"required"`
}

type UpdateUserRequest struct {
	UserRoleID   int    `json:"user_role_id" binding:"required"`
	UserStatusID int    `json:"user_status_id" binding:"required"`
	Nik          string `json:"nik" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
