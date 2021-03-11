package entities

type District struct {
	ID         int      `json:"id"`
	ProvinceID int      `json:"province_id"`
	Province   Province `json:"province"`
	Name       string   `json:"name"`
}
