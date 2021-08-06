package entities

import "time"

type Buyer struct {
	ID          int         `gorm:"not null" json:"id,omitempty"`
	UserID      int         `gorm:"not null" json:"user_id,omitempty"`
	User        *User       `json:"user,omitempty"`
	Nik         string      `gorm:"not null,unique" json:"nik,omitempty"`
	Name        string      `gorm:"not null" json:"name,omitempty"`
	Address     string      `gorm:"not null" json:"address,omitempty"`
	PhoneNumber string      `gorm:"not null" json:"phone_number,omitempty"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
	TpiID       int         `gorm:"index" json:"tpi_id,omitempty"`
	Tpi         *Tpi        `json:"tpi,omitempty"`
	BuyerTpi    []*BuyerTpi `json:"buyer_tpi,omitempty"`

	Status string `gorm:"-" json:"status,omitempty"`
}
