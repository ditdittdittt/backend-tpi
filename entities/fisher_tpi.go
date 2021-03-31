package entities

type FisherTpi struct {
	ID       int     `json:"id,omitempty"`
	FisherID int     `json:"fisher_id,omitempty"`
	Fisher   *Fisher `json:"fisher,omitempty"`
	TpiID    int     `json:"tpi_id,omitempty"`
	Tpi      *Tpi    `json:"tpi,omitempty"`
}
