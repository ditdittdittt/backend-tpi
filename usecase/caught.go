package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type CaughtUsecase interface {
	Create(caught *entities.Caught) error
	Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.CaughtItem, error)
	Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.CaughtItem, error)
	GetByID(id int) (entities.CaughtItem, error)
	Delete(id int) error
	Update(caught *entities.CaughtItem) error
}

type caughtUsecase struct {
	caughtRepository     mysql.CaughtRepository
	caughtItemRepository mysql.CaughtItemRepository
}

func (c *caughtUsecase) Update(caught *entities.CaughtItem) error {
	// insert log
	err := helper.InsertLog(caught.CaughtID, constant.Caught)
	if err != nil {
		return err
	}

	updateDateCaughtItem := map[string]interface{}{
		"fish_type_id": caught.FishTypeID,
		"weight":       caught.Weight,
		"weight_unit":  caught.WeightUnit,
	}

	updateDateCaught := map[string]interface{}{
		"user_id":         caught.Caught.UserID,
		"tpi_id":          caught.Caught.TpiID,
		"fisher_id":       caught.Caught.FisherID,
		"fishing_gear_id": caught.Caught.FishingGearID,
		"fishing_area_id": caught.Caught.FishingAreaID,
		"trip_day":        caught.Caught.TripDay,
	}

	err = c.caughtRepository.Update(caught.CaughtID, updateDateCaught)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Caught repository error")
	}

	err = c.caughtItemRepository.Update(caught.ID, updateDateCaughtItem)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Caught item repository error")
	}

	return nil
}

func (c *caughtUsecase) Delete(id int) error {
	caughtItem, err := c.caughtItemRepository.GetByID(id)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Caught item repository error")
	}

	caughtID := caughtItem.CaughtID

	err = c.caughtItemRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Caught item repository error")
	}

	caught, err := c.caughtRepository.GetByID(caughtID)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Caught repository error")
	}

	if len(caught.CaughtItem) == 0 {
		err = c.caughtRepository.Delete(caughtID)
		if err != nil {
			return stacktrace.Propagate(err, "[Delete] Caught repository error")
		}
	}

	return nil
}

func (c *caughtUsecase) GetByID(id int) (entities.CaughtItem, error) {
	caught, err := c.caughtItemRepository.GetByID(id)
	if err != nil {
		return entities.CaughtItem{}, stacktrace.Propagate(err, "[GetByID] Caught item repository error")
	}

	return caught, nil
}

func (c *caughtUsecase) Inquiry(fisherID int, fishTypeID int, tpiID int) ([]entities.CaughtItem, error) {
	queryMap := map[string]interface{}{
		"caughts.tpi_id":                tpiID,
		"caught_items.caught_status_id": 1,
	}

	if fisherID != 0 {
		queryMap["caughts.fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["caught_items.fish_type_id"] = fishTypeID
	}

	caughts, err := c.caughtItemRepository.Search(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Search] Caught item repository")
	}

	return caughts, nil
}

func (c *caughtUsecase) Index(fisherID int, fishTypeID int, caughtStatusID int, tpiID int) ([]entities.CaughtItem, error) {
	queryMap := map[string]interface{}{
		"caughts.tpi_id": tpiID,
	}

	if fisherID != 0 {
		queryMap["caughts.fisher_id"] = fisherID
	}

	if fishTypeID != 0 {
		queryMap["caught_items.fish_type_id"] = fishTypeID
	}

	if caughtStatusID != 0 {
		queryMap["caught_items.caught_status_id"] = caughtStatusID
	}

	date := time.Now().Format("2006-01-02")

	caughts, err := c.caughtItemRepository.Index(queryMap, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Caught item repository error")
	}

	return caughts, nil
}

func (c *caughtUsecase) Create(caught *entities.Caught) error {
	caught.CreatedAt = time.Now()
	caught.UpdatedAt = time.Now()

	currentDate := time.Now().Format("2006-01-02")

	existingCode, err := c.caughtRepository.GetLatestCode(currentDate)
	if err != nil {
		return stacktrace.Propagate(err, "[GetLatestCode] Caught repository error")
	}

	if existingCode != "" {
		latestID := existingCode[len(existingCode)-3:]
		intLatestID, _ := strconv.Atoi(latestID)
		intLatestID++
		caught.Code = formatDate(currentDate) + fmt.Sprintf("%03d", intLatestID)
	} else {
		caught.Code = formatDate(currentDate) + "001"
	}

	for index, caughtItem := range caught.CaughtItem {
		caughtItem.Code = caught.Code + fmt.Sprintf("%02d", index+1)
	}

	err = c.caughtRepository.Create(caught)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Caught repository error")
	}

	return nil
}

func NewCaughtUsecase(caughtRepository mysql.CaughtRepository, caughtItemRepository mysql.CaughtItemRepository) CaughtUsecase {
	return &caughtUsecase{caughtRepository: caughtRepository, caughtItemRepository: caughtItemRepository}
}

func formatDate(date string) string {
	result := strings.ReplaceAll(date, "-", "")
	return result[2:]
}
