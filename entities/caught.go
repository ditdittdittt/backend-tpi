package entities

import "time"

type Caught struct {
	ID             int          `json:"id"`
	UserID         int          `json:"user_id"`
	User           User         `json:"user"`
	TpiID          int          `json:"tpi_id"`
	Tpi            Tpi          `json:"tpi"`
	FisherID       int          `json:"fisher_id"`
	Fisher         Fisher       `json:"fisher"`
	FishType       FishType     `json:"fish_type"`
	FishingGearID  int          `json:"fishing_gear_id"`
	FishingGear    FishingGear  `json:"fishing_gear"`
	FishingAreaID  int          `json:"fishing_area_id"`
	FishingArea    FishingArea  `json:"fishing_area"`
	CaughtStatusID int          `json:"caught_status_id"`
	CaughtStatus   CaughtStatus `json:"caught_status"`
	CaughtData
	TripDay   int        `json:"trip_day"`
	SoldAt    *time.Time `json:"sold_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Code      string     `json:"code"`
}

type CaughtData struct {
	FishTypeID int     `json:"fish_type_id"`
	Weight     float64 `json:"weight"`
	WeightUnit string  `json:"weight_unit"`
}
