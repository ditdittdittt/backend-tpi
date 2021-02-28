package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type CaughtUsecase interface {
	Create(caught *entities.Caught, caughtData []entities.CaughtData) error
}

type caughtUsecase struct {
	caughtRepository repository.CaughtRepository
}

func (c *caughtUsecase) Create(caught *entities.Caught, caughtData []entities.CaughtData) error {

	caught.CaughtStatusID = 1
	caught.CreatedAt = time.Now()
	caught.UpdatedAt = time.Now()

	for _, caughtFish := range caughtData {
		caught.FishTypeID = caughtFish.FishTypeID
		caught.Weight = caughtFish.Weight
		caught.WeightUnit = caughtFish.WeightUnit

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
