package entities

import "time"

type Tpi struct {
	ID          int       `gorm:"not null" json:"id,omitempty"`
	DistrictID  int       `gorm:"not null" json:"district_id,omitempty"`
	District    *District `json:"district,omitempty"`
	UserID      int       `gorm:"not null" json:"user_id,omitempty"`
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Address     string    `gorm:"not null" json:"address,omitempty"`
	PhoneNumber string    `gorm:"not null" json:"phone_number,omitempty"`
	Pic         string    `gorm:"not null" json:"pic,omitempty"`
	Code        string    `gorm:"not null,unique" json:"code,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
