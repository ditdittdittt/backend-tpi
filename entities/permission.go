package entities

type Permission struct {
	ID   int    `gorm:"not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
