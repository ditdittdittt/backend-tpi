package entities

type FisherTpi struct {
	ID       int     `gorm:"not null" json:"id,omitempty"`
	FisherID int     `gorm:"not null" json:"fisher_id,omitempty"`
	Fisher   *Fisher `json:"fisher,omitempty"`
	TpiID    int     `gorm:"not null" json:"tpi_id,omitempty"`
	Tpi      *Tpi    `json:"tpi,omitempty"`
}
