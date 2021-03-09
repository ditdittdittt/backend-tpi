package entities

type TransactionItem struct {
	ID            int `json:"id"`
	AuctionID     int `json:"auction_id"`
	Auction       *Auction
	TransactionID int `json:"transaction_id"`
	Transaction   *Transaction
}
