package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishingGearUsecase interface {
	Create(fishingGear *entities.FishingGear) error
	Index() (fishingGears []entities.FishingGear, err error)
}

type fishingGearUsecase struct {
	FishingGearRepository repository.FishingGearRepository
}

func (f *fishingGearUsecase) Index() (fishingGears []entities.FishingGear, err error) {
	selectedField := []string{"code", "name"}

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
