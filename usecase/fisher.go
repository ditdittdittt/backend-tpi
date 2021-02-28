package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FisherUsecase interface {
	Create(fisher *entities.Fisher) error
}

type fisherUsecase struct {
	fisherRepository repository.FisherRepository
}

func (f *fisherUsecase) Create(fisher *entities.Fisher) error {

	fisher.CreatedAt = time.Now()
	fisher.UpdatedAt = time.Now()

	err := f.fisherRepository.Create(fisher)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Fisher repository err")
	}

	return nil
}

func NewFisherUsecase(fisherRepository repository.FisherRepository) FisherUsecase {
	return &fisherUsecase{fisherRepository: fisherRepository}
}
