package entities

import "time"

type Auction struct {
	ID        int       `json:"id,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	User      *User     `json:"user,omitempty"`
	TpiID     int       `json:"tpi_id,omitempty"`
	Tpi       *Tpi      `json:"tpi,omitempty"`
	CaughtID  int       `json:"caught_id,omitempty"`
	Caught    *Caught   `json:"caught,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Code      string    `json:"code,omitempty"`
}
