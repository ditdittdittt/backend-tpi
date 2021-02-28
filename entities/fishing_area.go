package entities

type FishingArea struct {
	ID                  int
	DistrictID          int
	District            District
	Name                string
	SouthLatitudeDegree string
	SouthLatitudeMinute string
	SouthLatitudeSecond string
	EastLongitudeDegree string
	EastLongitudeMinute string
	EastLongitudeSecond string
}
