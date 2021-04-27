package usecase

import (
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type FishTypeUsecase interface {
	Create(fishType *entities.FishType) error
	Index() (fishTypes []entities.FishType, err error)
	Delete(id int) error
	Update(fishType *entities.FishType) error
	GetByID(id int) (entities.FishType, error)
}

type fishTypeUsecase struct {
	fishTypeRepository mysql.FishTypeRepository
}

func (f *fishTypeUsecase) Delete(id int) error {
	err := f.fishTypeRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (f *fishTypeUsecase) Update(fishType *entities.FishType) error {
	err := f.fishTypeRepository.Update(fishType)
	if err != nil {
		return err
	}

	return nil
}

func (f *fishTypeUsecase) GetByID(id int) (entities.FishType, error) {
	fishType, err := f.fishTypeRepository.GetByID(id)
	if err != nil {
		return fishType, err
	}

	return fishType, nil
}

func (f *fishTypeUsecase) Index() (fishTypes []entities.FishType, err error) {
	selectedField := []string{"id", "code", "name"}

	fishTypes, err = f.fishTypeRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, err
	}

	return fishTypes, nil
}

func (f *fishTypeUsecase) Create(fishType *entities.FishType) error {

	err := f.fishTypeRepository.Create(fishType)
	if err != nil {
		return err
	}

	return nil
}

func NewFishTypeUsecase(fishTypeRepository mysql.FishTypeRepository) FishTypeUsecase {
	return &fishTypeUsecase{fishTypeRepository: fishTypeRepository}
}
