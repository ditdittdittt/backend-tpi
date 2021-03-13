package entities

import "time"

type Tpi struct {
	ID         int       `json:"id,omitempty"`
	DistrictID int       `json:"district_id,omitempty"`
	District   *District `json:"district,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	Code       string    `json:"code,omitempty"`
}
