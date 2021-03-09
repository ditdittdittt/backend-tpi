package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type CaughtUsecase interface {
	Create(caught *entities.Caught, caughtData []entities.CaughtData) error
	Index() ([]entities.Caught, error)
}

type caughtUsecase struct {
	caughtRepository repository.CaughtRepository
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

func NewCaughtUsecase(caughtRepository repository.CaughtRepository) CaughtUsecase {
	return &caughtUsecase{caughtRepository: caughtRepository}
}
