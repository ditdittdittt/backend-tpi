package entities

type FishingGear struct {
	ID         int      `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Code       string   `json:"code,omitempty"`
	DistrictID int      `json:"district_id,omitempty"`
	District   District `json:"district,omitempty"`
}
