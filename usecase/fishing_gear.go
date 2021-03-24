package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type FishingGearUsecase interface {
	Create(fishingGear *entities.FishingGear, tpiID int) error
	Index(tpiID int, districtID int) (fishingGears []entities.FishingGear, err error)
	Delete(id int) error
	Update(fishingGear *entities.FishingGear) error
	GetByID(id int) (entities.FishingGear, error)
}

type fishingGearUsecase struct {
	FishingGearRepository mysql.FishingGearRepository
	TpiRepository         mysql.TpiRepository
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

func (f *fishingGearUsecase) Index(tpiID int, districtID int) (fishingGears []entities.FishingGear, err error) {
	selectedField := []string{"fishing_gears.id", "fishing_gears.code", "fishing_gears.name", "fishing_gears.district_id"}

	if districtID == 0 {
		tpi, err := f.TpiRepository.GetByID(tpiID)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		districtID = tpi.DistrictID
	}

	queryMap := map[string]interface{}{
		"district_id": districtID,
	}

	fishingGears, err = f.FishingGearRepository.GetWithSelectedField(selectedField, queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fishing gear repository error")
	}

	return fishingGears, nil
}

func (f *fishingGearUsecase) Create(fishingGear *entities.FishingGear, tpiID int) error {
	if fishingGear.DistrictID == 0 {
		tpi, err := f.TpiRepository.GetByID(tpiID)
		if err != nil {
			return stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		fishingGear.DistrictID = tpi.DistrictID
	}

	err := f.FishingGearRepository.Create(fishingGear)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fishing Gear Repository err")
	}

	return nil
}

func NewFishingGearUsecase(fishingGearRepository mysql.FishingGearRepository, tpiRepository mysql.TpiRepository) FishingGearUsecase {
	return &fishingGearUsecase{FishingGearRepository: fishingGearRepository, TpiRepository: tpiRepository}
}
