package entities

import "time"

type Transaction struct {
	ID               int                `gorm:"not null" json:"id,omitempty"`
	UserID           int                `gorm:"not null" json:"user_id,omitempty"`
	User             *User              `json:"user,omitempty"`
	TpiID            int                `gorm:"not null, index" json:"tpi_id,omitempty"`
	Tpi              *Tpi               `json:"tpi,omitempty"`
	BuyerID          int                `gorm:"not null" json:"buyer_id,omitempty"`
	Buyer            *Buyer             `json:"buyer,omitempty"`
	DistributionArea string             `gorm:"not null" json:"distribution_area,omitempty"`
	Code             string             `gorm:"not null,unique" json:"code,omitempty"`
	TotalPrice       float64            `gorm:"not null" json:"total_price,omitempty"`
	CreatedAt        time.Time          `gorm:"not null, index" json:"created_at,omitempty"`
	UpdatedAt        time.Time          `gorm:"not null" json:"updated_at,omitempty"`
	TransactionItem  []*TransactionItem `gorm:"not null" json:"transaction_item,omitempty"`
}
