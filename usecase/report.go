package usecase

import (
	"fmt"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type ReportUsecase interface {
	GetProductionReport(tpiID int, from string, to string) (map[string]interface{}, error)
	GetTransactionReport(tpiID int, from string, to string) (map[string]interface{}, error)
}

type reportUsecase struct {
	caughtRepository          mysql.CaughtRepository
	auctionRepository         mysql.AuctionRepository
	transactionRepository     mysql.TransactionRepository
	fishTypeRepository        mysql.FishTypeRepository
	transactionItemRepository mysql.TransactionItemRepository
}

func (r *reportUsecase) GetTransactionReport(tpiID int, from string, to string) (map[string]interface{}, error) {
	data := map[string]interface{}{}

	transactionTotal, err := r.transactionRepository.GetTransactionTotal(tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTotalTransaction] Total transaction error")
	}

	data["transaction_total"] = transactionTotal

	weightTotal, err := r.caughtRepository.GetWeightTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
	}

	data["production_total"] = weightTotal

	productionValue, err := r.auctionRepository.GetPriceTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
	}

	data["production_value"] = productionValue

	averageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
	}

	data["transaction_speed"] = fmt.Sprintf("%.2f", averageTransactionSpeed/3600)

	permanentFisher, err := r.caughtRepository.GetFisherTotal("Tetap", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetFisherTotal] Caught repository error")
	}

	data["permanent_fisher"] = permanentFisher

	temporaryFisher, err := r.caughtRepository.GetFisherTotal("Pendatang", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetFisherTotal] Caught repository error")
	}

	data["temporary_fisher"] = temporaryFisher

	permanentBuyer, err := r.transactionRepository.GetBuyerTotal("Tetap", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTotalBuyer] Transaction repository error")
	}

	data["permanent_buyer"] = permanentBuyer

	temporaryBuyer, err := r.transactionRepository.GetBuyerTotal("Pendatang", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTotalBuyer] Transaction repository error")
	}

	data["temporary_buyer"] = temporaryBuyer

	transactionData, err := r.transactionItemRepository.GetReport(tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[Get] Transaction item repository error")
	}

	data["transaction_data"] = transactionData

	return data, nil
}

func (r *reportUsecase) GetProductionReport(tpiID int, from string, to string) (map[string]interface{}, error) {
	data := map[string]interface{}{}

	weightTotal, err := r.caughtRepository.GetWeightTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
	}

	data["production_total"] = weightTotal

	productionValue, err := r.auctionRepository.GetPriceTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
	}

	data["production_value"] = productionValue

	averageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
	}

	data["transaction_speed"] = fmt.Sprintf("%.2f", averageTransactionSpeed/3600)

	productionTable := make([]interface{}, 0)
	queryMap := []string{"id", "name", "code"}
	fishTypes, err := r.fishTypeRepository.GetWithSelectedField(queryMap)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetWithSelectedField] Fish type repository error")
	}

	for _, fishType := range fishTypes {
		fishWeightTotal, err := r.caughtRepository.GetWeightTotal(fishType.ID, tpiID, from, to)
		if err != nil {
			return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
		}

		fishProductionValue, err := r.auctionRepository.GetPriceTotal(fishType.ID, tpiID, from, to)
		if err != nil {
			return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
		}

		fishAverageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(fishType.ID, tpiID, from, to)
		if err != nil {
			return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
		}

		data := map[string]interface{}{
			"id":                fishType.ID,
			"code":              fishType.Code,
			"name":              fishType.Name,
			"production_total":  fishWeightTotal,
			"production_value":  fishProductionValue,
			"transaction_speed": fmt.Sprintf("%.2f", fishAverageTransactionSpeed/3600),
		}
		productionTable = append(productionTable, data)
	}

	data["production_table"] = productionTable
	return data, nil
}

func NewReportUsecase(
	caughtRepository mysql.CaughtRepository,
	auctionRepository mysql.AuctionRepository,
	transactionRepository mysql.TransactionRepository,
	fishTypeRepository mysql.FishTypeRepository,
	transactionItemRepository mysql.TransactionItemRepository) ReportUsecase {
	return &reportUsecase{
		caughtRepository:          caughtRepository,
		auctionRepository:         auctionRepository,
		transactionRepository:     transactionRepository,
		fishTypeRepository:        fishTypeRepository,
		transactionItemRepository: transactionItemRepository,
	}

}
