package entities

type FishingArea struct {
	ID                  int       `gorm:"not null" json:"id,omitempty"`
	DistrictID          int       `gorm:"not null" json:"district_id,omitempty"`
	District            *District `json:"district,omitempty"`
	Name                string    `gorm:"not null" gorm:"unique" json:"name,omitempty"`
	SouthLatitudeDegree string    `gorm:"not null" json:"south_latitude_degree,omitempty"`
	SouthLatitudeMinute string    `gorm:"not null" json:"south_latitude_minute,omitempty"`
	SouthLatitudeSecond string    `gorm:"not null" json:"south_latitude_second,omitempty"`
	EastLongitudeDegree string    `gorm:"not null" json:"east_longitude_degree,omitempty"`
	EastLongitudeMinute string    `gorm:"not null" json:"east_longitude_minute,omitempty"`
	EastLongitudeSecond string    `gorm:"not null" json:"east_longitude_second,omitempty"`
}
