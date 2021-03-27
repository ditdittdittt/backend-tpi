package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type FisherRepository interface {
	Create(fisher *entities.Fisher) error
	GetWithSelectedField(selectedField []string) (fishers []entities.Fisher, err error)
	GetByID(id int) (fisher entities.Fisher, err error)
	Update(fisher *entities.Fisher) error
	Delete(id int) error
	Get(query map[string]interface{}) (entities.Fisher, error)
	Index(query map[string]interface{}) ([]entities.Fisher, error)
}

type fisherRepository struct {
	db gorm.DB
}

func (f *fisherRepository) Index(query map[string]interface{}) ([]entities.Fisher, error) {
	var result []entities.Fisher

	err := f.db.Where(query).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (f *fisherRepository) Get(query map[string]interface{}) (entities.Fisher, error) {
	var result entities.Fisher

	err := f.db.Where(query).First(&result).Error
	if err != nil {
		return entities.Fisher{}, err
	}

	return result, nil
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
	return fisher, nil
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
	if fisher.TpiID == 0 {
		err := f.db.Omit("tpi_id").Create(&fisher).Error
		return err
	}
	err := f.db.Create(&fisher).Error
	return err
}

func NewFisherRepository(database gorm.DB) FisherRepository {
	return &fisherRepository{db: database}
}
