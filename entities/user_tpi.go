package entities

type UserTpi struct {
	ID     int  `json:"id,omitempty"`
	UserID int  `json:"user_id,omitempty"`
	User   User `json:"user,omitempty"`
	TpiID  int  `json:"tpi_id,omitempty"`
	Tpi    Tpi  `json:"tpi,omitempty"`
}
