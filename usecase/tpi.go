package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
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
		return nil, err
	}

	return tpis, nil
}

func (t *tpiUsecase) GetByID(id int) (entities.Tpi, error) {
	tpi, err := t.tpiRepository.GetByID(id)
	if err != nil {
		return entities.Tpi{}, err
	}

	return tpi, nil
}

func (t *tpiUsecase) Update(tpi *entities.Tpi) error {
	err := helper.InsertLog(tpi.ID, constant.Tpi)
	if err != nil {
		return err
	}

	err = t.tpiRepository.Update(tpi)
	if err != nil {
		return err
	}

	return nil
}

func (t *tpiUsecase) Delete(id int) error {
	err := t.tpiRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *tpiUsecase) Create(tpi *entities.Tpi) error {
	tpi.CreatedAt = time.Now()
	tpi.UpdatedAt = time.Now()

	existingCode, err := t.tpiRepository.GetLatestCode(tpi.DistrictID)
	if err != nil {
		return err
	}

	if existingCode != "" {
		latestID := existingCode[len(existingCode)-2:]
		intLatestID, _ := strconv.Atoi(latestID)
		intLatestID++
		tpi.Code = strconv.Itoa(tpi.DistrictID) + fmt.Sprintf("%02d", intLatestID)
	} else {
		tpi.Code = strconv.Itoa(tpi.DistrictID) + "01"
	}

	err = t.tpiRepository.Create(tpi)
	if err != nil {
		return err
	}

	return nil
}

func NewTpiUsecase(tpiRepository mysql.TpiRepository) TpiUsecase {
	return &tpiUsecase{tpiRepository: tpiRepository}
}
