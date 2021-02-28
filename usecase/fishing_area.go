package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishingAreaUsecase interface {
	Create(fishingArea *entities.FishingArea) error
}

type fishingAreaUsecase struct {
	fishingAreaRepository repository.FishingAreaRepository
}

func (f *fishingAreaUsecase) Create(fishingArea *entities.FishingArea) error {
	err := f.fishingAreaRepository.Create(fishingArea)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fishing Area Repository error")
	}

	return nil
}

func NewFishingAreaUsecase(fishingAreaRepository repository.FishingAreaRepository) FishingAreaUsecase {
	return &fishingAreaUsecase{fishingAreaRepository: fishingAreaRepository}
}
