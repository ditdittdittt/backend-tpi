package entities

import "time"

type Transaction struct {
	ID               int                `json:"id"`
	UserID           int                `json:"user_id"`
	User             *User              `json:"user"`
	TpiID            int                `json:"tpi_id"`
	Tpi              *Tpi               `json:"tpi"`
	BuyerID          int                `json:"buyer_id"`
	Buyer            *Buyer             `json:"buyer"`
	DistributionArea string             `json:"distribution_area"`
	Code             string             `json:"code"`
	TotalPrice       float64            `json:"total_price"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	TransactionItem  []*TransactionItem `json:"transaction_item"`
}
