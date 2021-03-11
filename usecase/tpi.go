package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type TpiUsecase interface {
	Create(tpi *entities.Tpi) error
}

type tpiUsecase struct {
	tpiRepository mysql.TpiRepository
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
