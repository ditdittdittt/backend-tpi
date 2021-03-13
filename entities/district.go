package entities

type District struct {
	ID         int       `json:"id,omitempty"`
	ProvinceID int       `json:"province_id,omitempty"`
	Province   *Province `json:"province,omitempty"`
	Name       string    `json:"name,omitempty"`
}
