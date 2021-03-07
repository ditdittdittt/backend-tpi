package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishTypeUsecase interface {
	Create(fishType *entities.FishType) error
	Index() (fishTypes []entities.FishType, err error)
	Delete(id int) error
	Update(fishType *entities.FishType) error
	GetByID(id int) (entities.FishType, error)
}

type fishTypeUsecase struct {
	fishTypeRepository repository.FishTypeRepository
}

func (f *fishTypeUsecase) Delete(id int) error {
	err := f.fishTypeRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Fish type repository error")
	}

	return nil
}

func (f *fishTypeUsecase) Update(fishType *entities.FishType) error {
	err := f.fishTypeRepository.Update(fishType)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Fish type repository error")
	}

	return nil
}

func (f *fishTypeUsecase) GetByID(id int) (entities.FishType, error) {
	fishType, err := f.fishTypeRepository.GetByID(id)
	if err != nil {
		return fishType, stacktrace.Propagate(err, "[GetByID] Fish type repository error")
	}

	return fishType, nil
}

func (f *fishTypeUsecase) Index() (fishTypes []entities.FishType, err error) {
	selectedField := []string{"id", "code", "name"}

	fishTypes, err = f.fishTypeRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fish type repository error")
	}

	return fishTypes, nil
}

func (f *fishTypeUsecase) Create(fishType *entities.FishType) error {

	err := f.fishTypeRepository.Create(fishType)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fish type repository")
	}

	return nil
}

func NewFishTypeUsecase(fishTypeRepository repository.FishTypeRepository) FishTypeUsecase {
	return &fishTypeUsecase{fishTypeRepository: fishTypeRepository}
}
