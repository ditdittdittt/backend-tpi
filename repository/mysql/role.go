package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type RoleRepository interface {
	Create(role *entities.Role) error
}

type roleRepository struct {
	db gorm.DB
}

func (r *roleRepository) Create(role *entities.Role) error {
	err := r.db.Create(&role).Error
	return err
}

func NewRoleRepository(database gorm.DB) RoleRepository {
	return &roleRepository{db: database}
}
