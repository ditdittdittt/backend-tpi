package entities

import "time"

type Buyer struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	User        *User     `json:"user,omitempty"`
	Nik         string    `json:"nik,omitempty"`
	Name        string    `json:"name,omitempty"`
	Address     string    `json:"address,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
