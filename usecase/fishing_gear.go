package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishingGearUsecase interface {
	Create(fishingGear *entities.FishingGear) error
	Index() (fishingGears []entities.FishingGear, err error)
	Delete(id int) error
	Update(fishingGear *entities.FishingGear) error
	GetByID(id int) (entities.FishingGear, error)
}

type fishingGearUsecase struct {
	FishingGearRepository repository.FishingGearRepository
}

func (f *fishingGearUsecase) Delete(id int) error {
	err := f.FishingGearRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Fishing gear repository error")
	}

	return nil
}

func (f *fishingGearUsecase) Update(fishingGear *entities.FishingGear) error {
	err := f.FishingGearRepository.Update(fishingGear)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Fishing gear repository error")
	}

	return nil
}

func (f *fishingGearUsecase) GetByID(id int) (entities.FishingGear, error) {
	fishingGear, err := f.FishingGearRepository.GetByID(id)
	if err != nil {
		return fishingGear, stacktrace.Propagate(err, "[GetByID] Fishing gear repository error")
	}

	return fishingGear, nil
}

func (f *fishingGearUsecase) Index() (fishingGears []entities.FishingGear, err error) {
	selectedField := []string{"id", "code", "name"}

	fishingGears, err = f.FishingGearRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fishing gear repository error")
	}

	return fishingGears, nil
}

func (f *fishingGearUsecase) Create(fishingGear *entities.FishingGear) error {
	err := f.FishingGearRepository.Create(fishingGear)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fishing Gear Repository err")
	}

	return nil
}

func NewFishingGearUsecase(fishingGearRepository repository.FishingGearRepository) FishingGearUsecase {
	return &fishingGearUsecase{FishingGearRepository: fishingGearRepository}
}
