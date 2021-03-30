package usecase

import (
	"time"

	"github.com/palantir/stacktrace"
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type FisherUsecase interface {
	Create(fisher *entities.Fisher, tpiID int, status string) error
	Delete(id int) error
	Update(fisher *entities.Fisher, status string) error
	GetByID(id int, tpiID int) (entities.Fisher, error)
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

func (f *fisherUsecase) GetByID(id int, tpiID int) (entities.Fisher, error) {
	fisher, err := f.fisherRepository.GetByID(id)
	if err != nil {
		return fisher, stacktrace.Propagate(err, "[GetByID] Fisher repository error")
	}

	if fisher.TpiID == tpiID {
		fisher.Status = constant.PermanentStatus
	} else {
		fisher.Status = constant.TemporaryStatus
	}

	return fisher, nil
}

func (f *fisherUsecase) Update(fisher *entities.Fisher, status string) error {
	// insert log
	err := helper.InsertLog(fisher.ID, constant.Fisher)
	if err != nil {
		return err
	}
	fisher.UpdatedAt = time.Now()

	updateData := map[string]interface{}{
		"user_id":      fisher.User,
		"nik":          fisher.Nik,
		"name":         fisher.Name,
		"nick_name":    fisher.NickName,
		"address":      fisher.Address,
		"ship_type":    fisher.ShipType,
		"abk_total":    fisher.AbkTotal,
		"phone_number": fisher.PhoneNumber,
	}

	existingFisher, err := f.fisherRepository.GetByID(fisher.ID)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Fisher repository error")
	}

	// Permanent to temporary
	if existingFisher.TpiID == fisher.TpiID && status == constant.TemporaryStatus {
		// remove tpi_id
		updateData["tpi_id"] = nil
		err = f.fisherRepository.Update(fisher.ID, updateData)
		if err != nil {
			return stacktrace.Propagate(err, "[Update] Fisher repository error")
		}

		// insert to fisher_tpis
		fisherTpi := &entities.FisherTpi{
			FisherID: fisher.ID,
			TpiID:    fisher.TpiID,
		}
		err = f.fisherTpiRepository.Create(fisherTpi)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Fisher tpi repository error")
		}
	}

	// Temporary to permanent
	if existingFisher.TpiID != fisher.TpiID && status == constant.PermanentStatus {
		// remove fisher_tpis
		err = f.fisherTpiRepository.Delete(map[string]interface{}{"fisher_id": fisher.ID, "tpi_id": fisher.TpiID})
		if err != nil {
			return stacktrace.Propagate(err, "[Delete] Fisher tpi repository error")
		}

		// update tpi_id
		updateData["tpi_id"] = fisher.TpiID
		err = f.fisherRepository.Update(fisher.ID, updateData)
		if err != nil {
			return stacktrace.Propagate(err, "[Update] Fisher repository error")
		}
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
		fisher.Status = constant.PermanentStatus
		result = append(result, fisher)
	}

	fisherTpis, err := f.fisherTpiRepository.Index(query)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Fisher tpi repository error")
	}
	for _, fisherTpi := range fisherTpis {
		fisherTpi.Fisher.Status = constant.TemporaryStatus
		result = append(result, *fisherTpi.Fisher)
	}

	return result, nil
}

func (f *fisherUsecase) Create(fisher *entities.Fisher, tpiID int, status string) error {
	fisher.CreatedAt = time.Now()
	fisher.UpdatedAt = time.Now()

	existingFisher, err := f.fisherRepository.Get(map[string]interface{}{"nik": fisher.Nik})
	if err != nil && err != gorm.ErrRecordNotFound {
		return stacktrace.Propagate(err, "[Get] Fisher repository error")
	}

	if err == gorm.ErrRecordNotFound {
		switch status {
		case constant.PermanentStatus:
			fisher.TpiID = tpiID
			err := f.fisherRepository.Create(fisher)
			if err != nil {
				return stacktrace.Propagate(err, "[Create] Fisher repository err")
			}

		case constant.TemporaryStatus:
			err = f.fisherRepository.Create(fisher)
			if err != nil {
				return stacktrace.Propagate(err, "[Create] Fisher repository err")
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

	if err == nil {
		err = helper.InsertLog(existingFisher.ID, constant.Fisher)
		if err != nil {
			return err
		}

		updateData := map[string]interface{}{
			"user_id":      fisher.User,
			"nik":          fisher.Nik,
			"name":         fisher.Name,
			"nick_name":    fisher.NickName,
			"address":      fisher.Address,
			"ship_type":    fisher.ShipType,
			"abk_total":    fisher.AbkTotal,
			"phone_number": fisher.PhoneNumber,
		}

		switch status {
		case constant.PermanentStatus:
			if existingFisher.TpiID == tpiID {
				err = f.fisherRepository.Update(existingFisher.ID, updateData)
				if err != nil {
					return err
				}
			}

			if existingFisher.TpiID != tpiID {
				updateData["tpi_id"] = tpiID

				err = f.fisherRepository.Update(existingFisher.ID, updateData)
				if err != nil {
					return err
				}

				err = f.fisherTpiRepository.Delete(map[string]interface{}{"fisher_id": existingFisher.ID, "tpi_id": tpiID})
				if err != nil {
					return err
				}
			}
		case constant.TemporaryStatus:
			if existingFisher.TpiID == tpiID {
				updateData["tpi_id"] = nil
				err = f.fisherRepository.Update(existingFisher.ID, updateData)
				if err != nil {
					return err
				}

				fisherTpi := &entities.FisherTpi{
					FisherID: existingFisher.ID,
					TpiID:    tpiID,
				}
				err = f.fisherTpiRepository.Create(fisherTpi)
				if err != nil {
					return err
				}
			}

			if existingFisher.TpiID != tpiID {
				fisherTpi := &entities.FisherTpi{
					FisherID: existingFisher.ID,
					TpiID:    tpiID,
				}
				err = f.fisherTpiRepository.Create(fisherTpi)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	return nil
}

func NewFisherUsecase(fisherRepository mysql.FisherRepository, fisherTpiRepository mysql.FisherTpiRepository) FisherUsecase {
	return &fisherUsecase{fisherRepository: fisherRepository, fisherTpiRepository: fisherTpiRepository}
}
