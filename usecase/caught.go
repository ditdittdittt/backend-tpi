package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type CaughtUsecase interface {
	Create(caught *entities.Caught, caughtData []entities.CaughtData) error
	Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.Caught, error)
	Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Caught, error)
	GetByID(id int) (entities.Caught, error)
	Delete(id int) error
	Update(caught *entities.Caught) error
}

type caughtUsecase struct {
	caughtRepository mysql.CaughtRepository
}

func (c *caughtUsecase) Update(caught *entities.Caught) error {
	data := map[string]interface{}{
		"fisher_id":       caught.FisherID,
		"fishing_gear_id": caught.FishingGearID,
		"fishing_area_id": caught.FishingAreaID,
		"trip_day":        caught.TripDay,
		"fish_type_id":    caught.FishTypeID,
		"weight":          caught.Weight,
		"weight_unit":     caught.WeightUnit,
	}

	err := c.caughtRepository.Update(caught, data)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Caught repository error")
	}

	return nil
}

func (c *caughtUsecase) Delete(id int) error {
	err := c.caughtRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Caught repository error")
	}

	return nil
}

func (c *caughtUsecase) GetByID(id int) (entities.Caught, error) {
	caught, err := c.caughtRepository.GetByID(id)
	if err != nil {
		return entities.Caught{}, stacktrace.Propagate(err, "[GetByID] Caught repository error")
	}

	return caught, nil
}

func (c *caughtUsecase) Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Caught, error) {
	queryMap := map[string]interface{}{
		"tpi_id":           tpiID,
		"caught_status_id": 1,
	}

	if fisherID != 0 {
		queryMap["fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["fish_type_id"] = fishTypeID
	}

	caughts, err := c.caughtRepository.Search(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Search] Caught repository")
	}

	return caughts, nil
}

func (c *caughtUsecase) Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.Caught, error) {
	queryMap := map[string]interface{}{
		"tpi_id": tpiID,
	}

	if fisherID != 0 {
		queryMap["fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["fish_type_id"] = fishTypeID
	}

	if caughtStatusID != 0 {
		queryMap["caught_status_id"] = caughtStatusID
	}

	startDate := time.Now().Format("2006-01-02")
	toDate := time.Now().String()

	caughts, err := c.caughtRepository.Get(queryMap, startDate, toDate)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Caught repository error")
	}

	return caughts, nil
}

func (c *caughtUsecase) Create(caught *entities.Caught, caughtData []entities.CaughtData) error {

	for _, caughtFish := range caughtData {
		caught := &entities.Caught{
			UserID:         caught.UserID,
			TpiID:          caught.TpiID,
			FisherID:       caught.FisherID,
			FishingGearID:  caught.FishingGearID,
			FishingAreaID:  caught.FishingAreaID,
			CaughtStatusID: 1,
			TripDay:        caught.TripDay,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			CaughtData:     caughtFish,
		}

		err := c.caughtRepository.Create(caught)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Caught repository error")
		}
	}

	return nil
}

func NewCaughtUsecase(caughtRepository mysql.CaughtRepository) CaughtUsecase {
	return &caughtUsecase{caughtRepository: caughtRepository}
}
