package entities

import "time"

type Transaction struct {
	ID               int
	UserID           int
	User             User
	TpiID            int
	Tpi              Tpi
	BuyerID          int
	Buyer            Buyer
	DistributionArea string
	Code             string
	TotalPrice       float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	TransactionItem  []TransactionItem
}
