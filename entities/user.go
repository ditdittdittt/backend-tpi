package entities

import "time"

type User struct {
	ID           int         `json:"id"`
	RoleID       int         `json:"role_id"`
	Role         *Role       `json:"role"`
	UserStatusID int         `json:"user_status_id"`
	UserStatus   *UserStatus `json:"user_status"`
	Nik          string      `json:"nik"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Username     string      `json:"username"`
	Password     string      `json:"-"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Token        string      `json:"token"`

	Permissions []string `gorm:"-" json:"permissions"`
}
