package entities

import "time"

type Auction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user"`
	TpiID     int       `json:"tpi_id"`
	Tpi       *Tpi      `json:"tpi"`
	CaughtID  int       `json:"caught_id"`
	Caught    *Caught   `json:"caught"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Code      string    `json:"code"`
}
