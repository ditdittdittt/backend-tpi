package entities

type UserSuperadmin struct {
	ID     int   `gorm:"not null" json:"id"`
	UserID int   `gorm:"not null" json:"user_id"`
	User   *User `gorm:"not null" json:"user"`
}
