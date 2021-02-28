package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type PermissionRepository interface {
	Create(permission *entities.Permission) error
}

type permissionRepository struct {
	db gorm.DB
}

func (p *permissionRepository) Create(permission *entities.Permission) error {
	err := p.db.Create(&permission).Error
	return err
}

func NewPermissionRepository(database gorm.DB) PermissionRepository {
	return &permissionRepository{db: database}
}