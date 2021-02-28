package usecase

import (
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FishTypeUsecase interface {
	Create(fishType *entities.FishType) error
}

type fishTypeUsecase struct {
	fishTypeRepository repository.FishTypeRepository
}

func (f *fishTypeUsecase) Create(fishType *entities.FishType) error {

	err := f.fishTypeRepository.Create(fishType)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fish type repository")
	}

	return nil
}

func NewFishTypeUsecase(fishTypeRepository repository.FishTypeRepository) FishTypeUsecase {
	return &fishTypeUsecase{fishTypeRepository: fishTypeRepository}
}
