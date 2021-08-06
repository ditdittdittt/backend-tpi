package entities

import "time"

type Caught struct {
	ID            int           `gorm:"not null" json:"id,omitempty"`
	UserID        int           `gorm:"not null" json:"user_id,omitempty"`
	User          *User         `json:"user,omitempty"`
	TpiID         int           `gorm:"not null, index" json:"tpi_id,omitempty"`
	Tpi           *Tpi          `json:"tpi,omitempty"`
	FisherID      int           `gorm:"not null, index" json:"fisher_id,omitempty"`
	Fisher        *Fisher       `json:"fisher,omitempty"`
	FishingGearID int           `gorm:"not null" json:"fishing_gear_id,omitempty"`
	FishingGear   *FishingGear  `json:"fishing_gear,omitempty"`
	FishingAreaID int           `gorm:"not null" json:"fishing_area_id,omitempty"`
	FishingArea   *FishingArea  `json:"fishing_area,omitempty"`
	TripDay       int           `gorm:"not null" json:"trip_day,omitempty"`
	CreatedAt     time.Time     `gorm:"not null, index" json:"created_at,omitempty"`
	UpdatedAt     time.Time     `gorm:"not null" json:"updated_at,omitempty"`
	Code          string        `gorm:"not null" gorm:"unique" json:"code,omitempty"`
	CaughtItem    []*CaughtItem `gorm:"not null" json:"caught_item,omitempty"`
}
