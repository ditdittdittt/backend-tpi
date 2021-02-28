package entities

import "time"

type Caught struct {
	ID		int
	UserID	int
	User	User
	TpiID	int
	Tpi		Tpi
	FisherID	int
	Fisher	Fisher
	FishTypeID	int
	FishType	FishType
	FishingGearID	int
	FishingGear	FishingGear
	FishingAreaID	int
	FishingArea	FishingArea
	CaughtStatusID	int
	CaughtStatus	CaughtStatus
	Weight			float64
	WeightUnit		string
	TripDay			int
	SoldAt			time.Time
	CreatedAt		time.Time
	UpdatedAt		time.Time
	Code			string
}