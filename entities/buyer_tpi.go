package entities

type BuyerTpi struct {
	ID      int    `json:"id,omitempty"`
	BuyerID int    `json:"buyer_id,omitempty"`
	Buyer   *Buyer `json:"buyer,omitempty"`
	TpiID   int    `json:"tpi_id,omitempty"`
	Tpi     *Tpi   `json:"tpi,omitempty"`
}
