package entities

import "time"

type Auction struct {
	ID           int         `json:"id,omitempty"`
	UserID       int         `json:"user_id,omitempty"`
	User         *User       `json:"user,omitempty"`
	TpiID        int         `json:"tpi_id,omitempty"`
	Tpi          *Tpi        `json:"tpi,omitempty"`
	CaughtItemID int         `json:"caught_item_id"`
	CaughtItem   *CaughtItem `json:"caught_item"`
	Price        float64     `json:"price,omitempty"`
	Code         string      `gorm:"unique" json:"code,omitempty"`
	CreatedAt    time.Time   `json:"created_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at,omitempty"`
}
