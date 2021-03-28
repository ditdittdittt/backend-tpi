package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type TpiUsecase interface {
	Create(tpi *entities.Tpi) error
	Index(districtID int) ([]entities.Tpi, error)
	GetByID(id int) (entities.Tpi, error)
	Update(tpi *entities.Tpi) error
	Delete(id int) error
}

type tpiUsecase struct {
	tpiRepository mysql.TpiRepository
}

func (t *tpiUsecase) Index(districtID int) ([]entities.Tpi, error) {
	queryMap := map[string]interface{}{}

	if districtID != 0 {
		queryMap["district_id"] = districtID
	}

	tpis, err := t.tpiRepository.Get(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Tpi repository error")
	}

	return tpis, nil
}

func (t *tpiUsecase) GetByID(id int) (entities.Tpi, error) {
	tpi, err := t.tpiRepository.GetByID(id)
	if err != nil {
		return entities.Tpi{}, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
	}

	return tpi, nil
}

func (t *tpiUsecase) Update(tpi *entities.Tpi) error {
	err := t.tpiRepository.Update(tpi)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Tpi repository error")
	}

	return nil
}

func (t *tpiUsecase) Delete(id int) error {
	err := t.tpiRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Tpi repository error")
	}

	return nil
}

func (t *tpiUsecase) Create(tpi *entities.Tpi) error {
	tpi.CreatedAt = time.Now()
	tpi.UpdatedAt = time.Now()

	err := t.tpiRepository.Create(tpi)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Tpi repository error")
	}

	return nil
}

func NewTpiUsecase(tpiRepository mysql.TpiRepository) TpiUsecase {
	return &tpiUsecase{tpiRepository: tpiRepository}
}
