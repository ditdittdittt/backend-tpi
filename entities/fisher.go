package entities

import "time"

type Fisher struct {
	ID          int          `gorm:"not null" json:"id,omitempty"`
	UserID      int          `gorm:"not null" json:"user_id,omitempty"`
	User        *User        `json:"user,omitempty"`
	Nik         string       `gorm:"not null,unique" json:"nik,omitempty"`
	Name        string       `gorm:"not null" json:"name,omitempty"`
	NickName    string       `gorm:"not null" json:"nick_name,omitempty"`
	Address     string       `gorm:"not null" json:"address,omitempty"`
	ShipType    string       `gorm:"not null" json:"ship_type,omitempty"`
	AbkTotal    int          `gorm:"not null" json:"abk_total,omitempty"`
	PhoneNumber string       `gorm:"not null" json:"phone_number,omitempty"`
	CreatedAt   time.Time    `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time    `gorm:"not null" json:"updated_at,omitempty"`
	TpiID       int          `gorm:"index" json:"tpi_id,omitempty"`
	Tpi         *Tpi         `json:"tpi,omitempty"`
	FisherTpi   []*FisherTpi `json:"fisher_tpi,omitempty"`

	Status string `gorm:"-" json:"status,omitempty"`
}
