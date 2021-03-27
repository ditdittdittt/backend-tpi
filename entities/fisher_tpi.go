package entities

type FisherTpi struct {
	ID       int     `json:"id"`
	FisherID int     `json:"fisher_id"`
	Fisher   *Fisher `json:"fisher"`
	TpiID    int     `json:"tpi_id"`
	Tpi      *Tpi    `json:"tpi"`
}
