package entities

type FishingArea struct {
	ID                  int       `json:"id,omitempty"`
	DistrictID          int       `json:"district_id,omitempty"`
	District            *District `json:"district,omitempty"`
	Name                string    `json:"name,omitempty"`
	SouthLatitudeDegree string    `json:"south_latitude_degree,omitempty"`
	SouthLatitudeMinute string    `json:"south_latitude_minute,omitempty"`
	SouthLatitudeSecond string    `json:"south_latitude_second,omitempty"`
	EastLongitudeDegree string    `json:"east_longitude_degree,omitempty"`
	EastLongitudeMinute string    `json:"east_longitude_minute,omitempty"`
	EastLongitudeSecond string    `json:"east_longitude_second,omitempty"`
}
