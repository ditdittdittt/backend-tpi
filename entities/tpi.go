package entities

import "time"

type Tpi struct {
	ID          int       `json:"id,omitempty"`
	DistrictID  int       `json:"district_id,omitempty"`
	District    *District `json:"district,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Address     string    `json:"address,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Pic         string    `json:"pic,omitempty"`
	Code        string    `gorm:"unique" json:"code,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
