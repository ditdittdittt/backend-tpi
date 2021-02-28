package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishingGearUsecase interface {
	Create(fishingGear *entities.FishingGear) error
}

type fishingGearUsecase struct {
	FishingGearRepository repository.FishingGearRepository
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
