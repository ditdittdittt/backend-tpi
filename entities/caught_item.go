package entities

type CaughtItem struct {
	ID             int           `gorm:"not null" json:"id"`
	CaughtID       int           `gorm:"not null, index" json:"caught_id,omitempty"`
	Caught         *Caught       `json:"caught,omitempty"`
	FishTypeID     int           `gorm:"not null, index" json:"fish_type_id,omitempty"`
	FishType       *FishType     `json:"fish_type,omitempty"`
	CaughtStatusID int           `gorm:"not null, index" json:"caught_status_id,omitempty"`
	CaughtStatus   *CaughtStatus `json:"caught_status,omitempty"`
	Weight         float64       `gorm:"not null" json:"weight,omitempty"`
	WeightUnit     string        `gorm:"not null" json:"weight_unit,omitempty"`
	Code           string        `gorm:"not null" gorm:"unique" json:"code,omitempty"`
}
