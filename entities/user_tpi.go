package entities

type UserTpi struct {
	ID     int   `gorm:"not null" json:"id,omitempty"`
	UserID int   `gorm:"not null" json:"user_id,omitempty"`
	User   *User `json:"user,omitempty"`
	TpiID  int   `gorm:"not null, index" json:"tpi_id,omitempty"`
	Tpi    *Tpi  `json:"tpi,omitempty"`
}
