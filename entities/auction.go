package entities

import "time"

type Auction struct {
	ID		int
	UserID	int
	User	User
	TpiID	int
	Tpi		Tpi
	CaughtID	int
	Caught	Caught
	BuyerID		int
	Buyer	Buyer
	DistributionArea	string
	Price		float64
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Code		string
}