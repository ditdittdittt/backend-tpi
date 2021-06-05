package entities

type TransactionItem struct {
	ID            int          `gorm:"not null" json:"id,omitempty"`
	AuctionID     int          `gorm:"not null" json:"auction_id,omitempty"`
	Auction       *Auction     `json:"auction,omitempty"`
	TransactionID int          `gorm:"not null" json:"transaction_id,omitempty"`
	Transaction   *Transaction `json:"transaction,omitempty"`
}
