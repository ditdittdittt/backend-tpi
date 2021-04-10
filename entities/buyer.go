package entities

import "time"

type Buyer struct {
	ID          int         `json:"id,omitempty"`
	UserID      int         `json:"user_id,omitempty"`
	User        *User       `json:"user,omitempty"`
	Nik         string      `gorm:"unique" json:"nik,omitempty"`
	Name        string      `json:"name,omitempty"`
	Address     string      `json:"address,omitempty"`
	PhoneNumber string      `json:"phone_number,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	TpiID       int         `json:"tpi_id,omitempty"`
	Tpi         *Tpi        `json:"tpi,omitempty"`
	BuyerTpi    []*BuyerTpi `json:"buyer_tpi,omitempty"`

	Status string `gorm:"-" json:"status,omitempty"`
}
