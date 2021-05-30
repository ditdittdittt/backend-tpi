package entities

type UserStatus struct {
	ID     int    `gorm:"not null" json:"id"`
	Status string `gorm:"not null" json:"status"`
}
