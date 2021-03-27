package usecase

import (
	"time"

	"github.com/palantir/stacktrace"
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type FisherUsecase interface {
	Create(fisher *entities.Fisher, tpiID int, status string) error
	Delete(id int) error
	Update(fisher *entities.Fisher) error
	GetByID(id int) (entities.Fisher, error)
	Index(tpiID int) (fishers []entities.Fisher, err error)
}

type fisherUsecase struct {
	fisherRepository    mysql.FisherRepository
	fisherTpiRepository mysql.FisherTpiRepository
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

func (f *fisherUsecase) Index(tpiID int) (fishers []entities.Fisher, err error) {
	result := make([]entities.Fisher, 0)

	query := map[string]interface{}{
		"tpi_id": tpiID,
	}

	fishers, err = f.fisherRepository.Index(query)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Fisher repository error")
	}
	for _, fisher := range fishers {
		fisher.Status = "Tetap"
		result = append(result, fisher)
	}

	fisherTpis, err := f.fisherTpiRepository.Index(query)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Fisher tpi repository error")
	}
	for _, fisherTpi := range fisherTpis {
		fisherTpi.Fisher.Status = "Pendatang"
		result = append(result, *fisherTpi.Fisher)
	}

	return result, nil
}

func (f *fisherUsecase) Create(fisher *entities.Fisher, tpiID int, status string) error {
	fisher.CreatedAt = time.Now()
	fisher.UpdatedAt = time.Now()

	switch status {
	case "Tetap":
		fisher.TpiID = tpiID

		err := f.fisherRepository.Create(fisher)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Fisher repository err")
		}

	case "Pendatang":
		existedFisher, err := f.fisherRepository.Get(map[string]interface{}{"nik": fisher.Nik})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = f.fisherRepository.Create(fisher)
				if err != nil {
					return stacktrace.Propagate(err, "[Create] Fisher repository err")
				}
			} else {
				return stacktrace.Propagate(err, "[Get] Fisher repository error")
			}
		}

		if existedFisher.Nik == fisher.Nik {
			fisher.ID = existedFisher.ID
		}

		fisherTpi := &entities.FisherTpi{
			FisherID: fisher.ID,
			TpiID:    tpiID,
		}

		err = f.fisherTpiRepository.Create(fisherTpi)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Fisher tpi repository")
		}
	}

	return nil
}

func NewFisherUsecase(fisherRepository mysql.FisherRepository, fisherTpiRepository mysql.FisherTpiRepository) FisherUsecase {
	return &fisherUsecase{fisherRepository: fisherRepository, fisherTpiRepository: fisherTpiRepository}
}
