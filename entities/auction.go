package entities

import "time"

type Auction struct {
	ID           int         `gorm:"not null" json:"id,omitempty"`
	UserID       int         `gorm:"not null" json:"user_id,omitempty"`
	User         *User       `json:"user,omitempty"`
	TpiID        int         `gorm:"not null, index" json:"tpi_id,omitempty"`
	Tpi          *Tpi        `json:"tpi,omitempty"`
	CaughtItemID int         `gorm:"not null, index" json:"caught_item_id"`
	CaughtItem   *CaughtItem `json:"caught_item"`
	Price        float64     `gorm:"not null" json:"price,omitempty"`
	Code         string      `gorm:"not null,unique" json:"code,omitempty"`
	CreatedAt    time.Time   `gorm:"not null, index" json:"created_at,omitempty"`
	UpdatedAt    time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}
