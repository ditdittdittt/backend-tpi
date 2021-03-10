package entities

import "time"

type Caught struct {
	ID             int           `json:"id,omitempty"`
	UserID         int           `json:"user_id,omitempty"`
	User           *User         `json:"user,omitempty"`
	TpiID          int           `json:"tpi_id,omitempty"`
	Tpi            *Tpi          `json:"tpi,omitempty"`
	FisherID       int           `json:"fisher_id,omitempty"`
	Fisher         *Fisher       `json:"fisher,omitempty"`
	FishType       *FishType     `json:"fish_type,omitempty"`
	FishingGearID  int           `json:"fishing_gear_id,omitempty"`
	FishingGear    *FishingGear  `json:"fishing_gear,omitempty"`
	FishingAreaID  int           `json:"fishing_area_id,omitempty"`
	FishingArea    *FishingArea  `json:"fishing_area,omitempty"`
	CaughtStatusID int           `json:"caught_status_id,omitempty"`
	CaughtStatus   *CaughtStatus `json:"caught_status,omitempty"`
	CaughtData
	TripDay   int       `json:"trip_day,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Code      string    `json:"code,omitempty"`
}

type CaughtData struct {
	FishTypeID int     `json:"fish_type_id,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	WeightUnit string  `json:"weight_unit,omitempty"`
}
