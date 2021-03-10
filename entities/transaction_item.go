package entities

type TransactionItem struct {
	ID            int          `json:"id,omitempty"`
	AuctionID     int          `json:"auction_id,omitempty"`
	Auction       *Auction     `json:"auction,omitempty"`
	TransactionID int          `json:"transaction_id,omitempty"`
	Transaction   *Transaction `json:"transaction,omitempty"`
}
