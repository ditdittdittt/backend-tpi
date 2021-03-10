package entities

import "time"

type Transaction struct {
	ID               int                `json:"id,omitempty"`
	UserID           int                `json:"user_id,omitempty"`
	User             *User              `json:"user,omitempty"`
	TpiID            int                `json:"tpi_id,omitempty"`
	Tpi              *Tpi               `json:"tpi,omitempty"`
	BuyerID          int                `json:"buyer_id,omitempty"`
	Buyer            *Buyer             `json:"buyer,omitempty"`
	DistributionArea string             `json:"distribution_area,omitempty"`
	Code             string             `json:"code,omitempty"`
	TotalPrice       float64            `json:"total_price,omitempty"`
	CreatedAt        time.Time          `json:"created_at,omitempty"`
	UpdatedAt        time.Time          `json:"updated_at,omitempty"`
	TransactionItem  []*TransactionItem `json:"transaction_item,omitempty"`
}
