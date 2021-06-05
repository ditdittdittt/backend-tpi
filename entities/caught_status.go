package entities

type CaughtStatus struct {
	ID     int    `gorm:"not null" json:"id,omitempty"`
	Status string `gorm:"not null" json:"status,omitempty"`
}
