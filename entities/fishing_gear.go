package entities

type FishingGear struct {
	ID         int       `gorm:"not null" json:"id,omitempty"`
	Name       string    `gorm:"not null" json:"name,omitempty"`
	Code       string    `gorm:"not null, unique" json:"code,omitempty"`
	DistrictID int       `gorm:"not null" json:"district_id,omitempty"`
	District   *District `json:"district,omitempty"`
}
