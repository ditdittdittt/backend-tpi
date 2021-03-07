package repository

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FisherRepository interface {
	Create(fisher *entities.Fisher) error
	GetByID(id int) (fisher entities.Fisher, err error)
	GetWithSelectedField(selectedField []string) (fishers []entities.Fisher, err error)
	Update(fisher *entities.Fisher) error
	Delete(id int) error
}

type fisherRepository struct {
	db gorm.DB
}

func (f *fisherRepository) Delete(id int) error {
	err := f.db.Delete(&entities.Fisher{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (f *fisherRepository) GetByID(id int) (fisher entities.Fisher, err error) {
	err = f.db.First(&fisher, id).Error
	if err != nil {
		return entities.Fisher{}, err
	}
	return fisher, err
}

func (f *fisherRepository) Update(fisher *entities.Fisher) error {
	err := f.db.Model(&fisher).Updates(fisher).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *fisherRepository) GetWithSelectedField(selectedField []string) (fishers []entities.Fisher, err error) {
	err = f.db.Select(selectedField).Find(&fishers).Error
	if err != nil {
		return nil, err
	}
	return fishers, err
}

func (f *fisherRepository) Create(fisher *entities.Fisher) error {
	err := f.db.Create(&fisher).Error
	return err
}

func NewFisherRepository(database gorm.DB) FisherRepository {
	return &fisherRepository{db: database}
}
