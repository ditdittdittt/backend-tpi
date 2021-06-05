package entities

type District struct {
	ID         int       `gorm:"not null" json:"id,omitempty"`
	ProvinceID int       `gorm:"not null" json:"province_id,omitempty"`
	Province   *Province `json:"province,omitempty"`
	Name       string    `gorm:"not null" json:"name,omitempty"`
}
