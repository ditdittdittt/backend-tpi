package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishingAreaUsecase interface {
	Create(fishingArea *entities.FishingArea) error
	Index() (fishingAreas []entities.FishingArea, err error)
}

type fishingAreaUsecase struct {
	fishingAreaRepository repository.FishingAreaRepository
}

func (f *fishingAreaUsecase) Index() (fishingAreas []entities.FishingArea, err error) {
	selectedField := []string{"name", "south_latitude_degree", "south_latitude_minute", "south_latitude_second", "east_longitude_degree", "east_longitude_minute", "east_longitude_second"}

	fishingAreas, err = f.fishingAreaRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fishing area repository error")
	}

	return fishingAreas, nil
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
