package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type FisherUsecase interface {
	Create(fisher *entities.Fisher) error
	Delete(id int) error
	Update(fisher *entities.Fisher) error
	GetByID(id int) (entities.Fisher, error)
	Index() (fishers []entities.Fisher, err error)
}

type fisherUsecase struct {
	fisherRepository repository.FisherRepository
}

func (f *fisherUsecase) Delete(id int) error {
	err := f.fisherRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Fisher repository error")
	}

	return nil
}

func (f *fisherUsecase) GetByID(id int) (entities.Fisher, error) {
	fisher, err := f.fisherRepository.GetByID(id)
	if err != nil {
		return fisher, stacktrace.Propagate(err, "[GetByID] Fisher repository error")
	}

	return fisher, nil
}

func (f *fisherUsecase) Update(fisher *entities.Fisher) error {
	fisher.UpdatedAt = time.Now()

	err := f.fisherRepository.Update(fisher)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Fisher repository error")
	}

	return nil
}

func (f *fisherUsecase) Index() (fishers []entities.Fisher, err error) {
	selectedField := []string{"id", "nik", "name", "status", "address", "phone_number", "ship_type", "abk_total", "created_at", "updated_at"}

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
