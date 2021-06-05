package entities

type Role struct {
	ID         int           `gorm:"not null" json:"id"`
	Name       string        `gorm:"not null" json:"name"`
	Permission []*Permission `gorm:"many2many:role_permissions" json:"-"`
}
