package entities

type Role struct {
	ID		int
	Name	string
	Permission	[]Permission	`gorm:"many2many:role_permissions"`
}