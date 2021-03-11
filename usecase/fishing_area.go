package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type FishingAreaUsecase interface {
	Create(fishingArea *entities.FishingArea, tpiID int) error
	Delete(id int) error
	Update(fishingArea *entities.FishingArea) error
	GetByID(id int) (entities.FishingArea, error)
	Index() (fishingAreas []entities.FishingArea, err error)
}

type fishingAreaUsecase struct {
	fishingAreaRepository mysql.FishingAreaRepository
	tpiRepository         mysql.TpiRepository
}

func (f *fishingAreaUsecase) Delete(id int) error {
	err := f.fishingAreaRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Fishing area repository error")
	}

	return nil
}

func (f *fishingAreaUsecase) Update(fishingArea *entities.FishingArea) error {
	err := f.fishingAreaRepository.Update(fishingArea)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Fishing area repository error")
	}

	return nil
}

func (f *fishingAreaUsecase) GetByID(id int) (entities.FishingArea, error) {
	fishingArea, err := f.fishingAreaRepository.GetByID(id)
	if err != nil {
		return fishingArea, stacktrace.Propagate(err, "[GetByID] Fishing area repository error")
	}

	return fishingArea, nil
}

func (f *fishingAreaUsecase) Index() (fishingAreas []entities.FishingArea, err error) {
	selectedField := []string{
		"fishing_areas.id",
		"fishing_areas.name",
		"fishing_areas.south_latitude_degree",
		"fishing_areas.south_latitude_minute",
		"fishing_areas.south_latitude_second",
		"fishing_areas.east_longitude_degree",
		"fishing_areas.east_longitude_minute",
		"fishing_areas.east_longitude_second",
	}

	fishingAreas, err = f.fishingAreaRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fishing area repository error")
	}

	return fishingAreas, nil
}

func (f *fishingAreaUsecase) Create(fishingArea *entities.FishingArea, tpiID int) error {
	if fishingArea.DistrictID == 0 {
		tpi, err := f.tpiRepository.GetByID(tpiID)
		if err != nil {
			return stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		fishingArea.DistrictID = tpi.DistrictID
	}

	err := f.fishingAreaRepository.Create(fishingArea)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fishing Area Repository error")
	}

	return nil
}

func NewFishingAreaUsecase(fishingAreaRepository mysql.FishingAreaRepository, tpiRepository mysql.TpiRepository) FishingAreaUsecase {
	return &fishingAreaUsecase{fishingAreaRepository: fishingAreaRepository, tpiRepository: tpiRepository}
}
