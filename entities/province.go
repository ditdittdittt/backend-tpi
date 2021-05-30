package entities

type Province struct {
	ID   int    `gorm:"not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
