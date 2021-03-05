package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FisherUsecase interface {
	Create(fisher *entities.Fisher) error
	Index() (fishers []entities.Fisher, err error)
}

type fisherUsecase struct {
	fisherRepository repository.FisherRepository
}

func (f *fisherUsecase) Index() (fishers []entities.Fisher, err error) {
	selectedField := []string{"nik", "name", "status", "address", "phone_number", "ship_type", "abk_total", "created_at", "updated_at"}

	fishers, err = f.fisherRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Fisher repository error")
	}

	return fishers, nil
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
