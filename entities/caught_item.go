package entities

type CaughtItem struct {
	ID             int           `json:"id"`
	CaughtID       int           `json:"caught_id,omitempty"`
	Caught         *Caught       `json:"caught,omitempty"`
	FishTypeID     int           `json:"fish_type_id,omitempty"`
	FishType       *FishType     `json:"fish_type,omitempty"`
	CaughtStatusID int           `json:"caught_status_id,omitempty"`
	CaughtStatus   *CaughtStatus `json:"caught_status,omitempty"`
	Weight         float64       `json:"weight,omitempty"`
	WeightUnit     string        `json:"weight_unit,omitempty"`
	Code           string        `json:"code,omitempty"`
}
