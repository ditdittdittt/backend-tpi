package entities

type FishType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `gorm:"unique" json:"code"`
}
