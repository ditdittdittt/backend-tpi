package entities

type BuyerTpi struct {
	ID      int    `json:"id"`
	BuyerID int    `json:"buyer_id"`
	Buyer   *Buyer `json:"buyer"`
	TpiID   int    `json:"tpi_id"`
	Tpi     *Tpi   `json:"tpi"`
}
