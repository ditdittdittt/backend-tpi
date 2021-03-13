package entities

type Role struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Permission []*Permission `gorm:"many2many:role_permissions" json:"-"`
}
