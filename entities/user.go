package entities

import "time"

type User struct {
	ID           int         `gorm:"not null" json:"id,omitempty"`
	RoleID       int         `gorm:"not null" json:"role_id,omitempty"`
	Role         *Role       `json:"role,omitempty"`
	UserStatusID int         `gorm:"not null" json:"user_status_id,omitempty"`
	UserStatus   *UserStatus `json:"user_status,omitempty"`
	Nik          string      `gorm:"not null" gorm:"unique" json:"nik,omitempty"`
	Name         string      `gorm:"not null" json:"name,omitempty"`
	Address      string      `gorm:"not null" json:"address,omitempty"`
	Username     string      `gorm:"not null,unique" json:"username,omitempty"`
	Password     string      `gorm:"not null" json:"-"`
	CreatedAt    time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time   `gorm:"not null" json:"updated_at,omitempty"`

	Permissions  []string `gorm:"-" json:"permissions,omitempty"`
	LocationID   int      `gorm:"-" json:"location_id,omitempty"`
	LocationName string   `gorm:"-" json:"location_name,omitempty"`
}
