package entities

import "time"

type Caught struct {
	ID             int
	UserID         int
	User           User
	TpiID          int
	Tpi            Tpi
	FisherID       int
	Fisher         Fisher
	FishType       FishType
	FishingGearID  int
	FishingGear    FishingGear
	FishingAreaID  int
	FishingArea    FishingArea
	CaughtStatusID int
	CaughtStatus   CaughtStatus
	CaughtData
	TripDay   int
	SoldAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Code      string
}

type CaughtData struct {
	FishTypeID int     `json:"fish_type_id"`
	Weight     float64 `json:"weight"`
	WeightUnit string  `json:"weight_unit"`
}
