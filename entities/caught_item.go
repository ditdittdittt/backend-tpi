package entities

type CaughtItem struct {
	ID             int           `json:"id"`
	CaughtID       int           `json:"caught_id"`
	Caught         *Caught       `json:"caught"`
	FishTypeID     int           `json:"fish_type_id"`
	FishType       *FishType     `json:"fish_type"`
	CaughtStatusID int           `json:"caught_status_id,omitempty"`
	CaughtStatus   *CaughtStatus `json:"caught_status,omitempty"`
	Weight         float64       `json:"weight"`
	WeightUnit     string        `json:"weight_unit"`
	Code           string        `json:"code"`
}
