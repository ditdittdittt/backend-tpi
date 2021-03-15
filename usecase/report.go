package usecase

import (
	"fmt"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type ReportUsecase interface {
	Get(tpiID int, from string, to string) (map[string]interface{}, error)
}

type reportUsecase struct {
	caughtRepository      mysql.CaughtRepository
	auctionRepository     mysql.AuctionRepository
	transactionRepository mysql.TransactionRepository
	fishTypeRepository    mysql.FishTypeRepository
}

func (r *reportUsecase) Get(tpiID int, from string, to string) (map[string]interface{}, error) {
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

func NewReportUsecase(caughtRepository mysql.CaughtRepository, auctionRepository mysql.AuctionRepository, transactionRepository mysql.TransactionRepository, fishTypeRepository mysql.FishTypeRepository) ReportUsecase {
	return &reportUsecase{
		caughtRepository:      caughtRepository,
		auctionRepository:     auctionRepository,
		transactionRepository: transactionRepository,
		fishTypeRepository:    fishTypeRepository,
	}

}
