package entities

import "time"

type User struct {
	ID           int         `json:"id,omitempty"`
	RoleID       int         `json:"role_id,omitempty"`
	Role         *Role       `json:"role,omitempty"`
	UserStatusID int         `json:"user_status_id,omitempty"`
	UserStatus   *UserStatus `json:"user_status,omitempty"`
	Nik          string      `json:"nik,omitempty"`
	Name         string      `json:"name,omitempty"`
	Address      string      `json:"address,omitempty"`
	Username     string      `json:"username,omitempty"`
	Password     string      `json:"-"`
	CreatedAt    time.Time   `json:"created_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at,omitempty"`
	Token        string      `json:"token,omitempty"`

	Permissions []string `gorm:"-" json:"permissions,omitempty"`
}
