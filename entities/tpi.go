package entities

import "time"

type Tpi struct {
	ID         int `json:"id"`
	DistrictID int
	District   District
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Code       string
}
