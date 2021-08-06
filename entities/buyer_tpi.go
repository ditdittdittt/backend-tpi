package entities

type BuyerTpi struct {
	ID      int    `gorm:"not null" json:"id,omitempty"`
	BuyerID int    `gorm:"not null" json:"buyer_id,omitempty"`
	Buyer   *Buyer `json:"buyer,omitempty"`
	TpiID   int    `gorm:"not null, index" json:"tpi_id,omitempty"`
	Tpi     *Tpi   `json:"tpi,omitempty"`
}
