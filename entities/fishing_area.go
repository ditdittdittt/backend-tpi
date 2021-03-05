package entities

type FishingArea struct {
	ID                  int       `json:"id"`
	DistrictID          int       `json:"district_id"`
	District            *District `json:"district"`
	Name                string    `json:"name"`
	SouthLatitudeDegree string    `json:"south_latitude_degree"`
	SouthLatitudeMinute string    `json:"south_latitude_minute"`
	SouthLatitudeSecond string    `json:"south_latitude_second"`
	EastLongitudeDegree string    `json:"east_longitude_degree"`
	EastLongitudeMinute string    `json:"east_longitude_minute"`
	EastLongitudeSecond string    `json:"east_longitude_second"`
}
