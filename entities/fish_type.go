package entities

type FishType struct {
	ID   int    `gorm:"not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Code string `gorm:"not null" gorm:"unique" json:"code"`
}
