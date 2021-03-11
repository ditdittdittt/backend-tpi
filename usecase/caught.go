package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type CaughtUsecase interface {
	Create(caught *entities.Caught, caughtData []entities.CaughtData) error
	Index() ([]entities.Caught, error)
	Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.Caught, error)
}

type caughtUsecase struct {
	caughtRepository mysql.CaughtRepository
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

func (c *caughtUsecase) Index() ([]entities.Caught, error) {
	caughts, err := c.caughtRepository.Get()
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
