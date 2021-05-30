package entities

type UserDistrict struct {
	ID         int       `gorm:"not null" json:"id"`
	UserID     int       `gorm:"not null" json:"user_id"`
	User       *User     `json:"user"`
	DistrictID int       `gorm:"not null" json:"district_id"`
	District   *District `json:"district"`
}
