package usecase

import (
	"fmt"
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type DashboardUsecase interface {
	GetFisherAndBuyerTotal(tpiID int, districtID int) (map[string]interface{}, error)
	GetDashboardDetail(tpiID int, districtID int, queryType string) (map[string]interface{}, error)
}

type dashboardUsecase struct {
	caughtRepository      mysql.CaughtRepository
	auctionRepository     mysql.AuctionRepository
	transactionRepository mysql.TransactionRepository
	tpiRepository         mysql.TpiRepository
}

var queryTypeMap = map[string]bool{
	"daily":   true,
	"monthly": true,
	"yearly":  true,
}

func (d *dashboardUsecase) GetFisherAndBuyerTotal(tpiID int, districtID int) (map[string]interface{}, error) {
	if districtID != 0 {
		tpis, err := d.tpiRepository.Get(map[string]interface{}{"district_id": districtID})
		if err != nil {
			return nil, stacktrace.Propagate(err, "[Get] Tpi repository error")
		}

		permanentFisherTotal := 0
		permanentBuyerTotal := 0
		temporaryFisherTotal := 0
		temporaryBuyerTotal := 0

		for _, tpi := range tpis {
			permanentFisherTotalPerTpi, err := d.caughtRepository.GetFisherTotalDashboard(tpi.ID, constant.PermanentStatus)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetFisherTotalDashboard] Caught repository error")
			}

			temporaryFisherTotalPerTpi, err := d.caughtRepository.GetFisherTotalDashboard(tpi.ID, constant.TemporaryStatus)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetFisherTotalDashboard] Caught repository error")
			}

			permanentBuyerTotalPerTpi, err := d.transactionRepository.GetBuyerTotalDashboard(tpi.ID, constant.PermanentStatus)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetBuyerTotalDashboard] Transaction repository error")
			}

			temporaryBuyerTotalPerTpi, err := d.transactionRepository.GetBuyerTotalDashboard(tpi.ID, constant.TemporaryStatus)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetBuyerTotalDashboard] Transaction repository error")
			}

			permanentFisherTotal += permanentFisherTotalPerTpi
			permanentBuyerTotal += permanentBuyerTotalPerTpi
			temporaryFisherTotal += temporaryFisherTotalPerTpi
			temporaryBuyerTotal += temporaryBuyerTotalPerTpi
		}

		result := map[string]interface{}{
			"fisher_total":           permanentFisherTotal + temporaryFisherTotal,
			"buyer_total":            permanentBuyerTotal + temporaryBuyerTotal,
			"permanent_fisher_total": permanentFisherTotal,
			"permanent_buyer_total":  permanentBuyerTotal,
			"temporary_fisher_total": temporaryFisherTotal,
			"temporary_buyer_total":  temporaryBuyerTotal,
		}

		return result, nil
	}

	permanentFisherTotal, err := d.caughtRepository.GetFisherTotalDashboard(tpiID, constant.PermanentStatus)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetFisherTotalDashboard] Caught repository error")
	}

	temporaryFisherTotal, err := d.caughtRepository.GetFisherTotalDashboard(tpiID, constant.TemporaryStatus)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetFisherTotalDashboard] Caught repository error")
	}

	permanentBuyerTotal, err := d.transactionRepository.GetBuyerTotalDashboard(tpiID, constant.PermanentStatus)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetBuyerTotalDashboard] Transaction repository error")
	}

	temporaryBuyerTotal, err := d.transactionRepository.GetBuyerTotalDashboard(tpiID, constant.TemporaryStatus)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetBuyerTotalDashboard] Transaction repository error")
	}

	result := map[string]interface{}{
		"fisher_total":           permanentFisherTotal + temporaryFisherTotal,
		"buyer_total":            permanentBuyerTotal + temporaryBuyerTotal,
		"permanent_fisher_total": permanentFisherTotal,
		"permanent_buyer_total":  permanentBuyerTotal,
		"temporary_fisher_total": temporaryFisherTotal,
		"temporary_buyer_total":  temporaryBuyerTotal,
	}

	return result, nil
}

func (d *dashboardUsecase) GetDashboardDetail(tpiID int, districtID int, queryType string) (map[string]interface{}, error) {
	if !queryTypeMap[queryType] {
		return nil, stacktrace.NewError("Error invalid query type")
	}

	date := time.Now().Format("2006-01-02")

	if districtID != 0 {
		tpis, err := d.tpiRepository.Get(map[string]interface{}{"district_id": districtID})
		if err != nil {
			return nil, stacktrace.Propagate(err, "[Get] Tpi repository error")
		}

		productionTotal := 0.0
		transactionTotal := 0.0
		transactionSpeed := 0.0

		for _, tpi := range tpis {
			productionTotalPerTpi, err := d.caughtRepository.GetProductionTotalDashboard(tpi.ID, queryType, date)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetProductionTotalDashboard] Caught repository error")
			}

			transactionTotalPerTpi, err := d.auctionRepository.GetTransactionValueDashboard(tpi.ID, queryType, date)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetTransactionTotalDashboard] Auction repository error")
			}

			transactionSpeedPerTpi, err := d.auctionRepository.GetTransactionSpeedDashboard(tpi.ID, queryType, date)
			if err != nil {
				return nil, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
			}

			productionTotal += productionTotalPerTpi
			transactionTotal += transactionTotalPerTpi
			transactionSpeed += transactionSpeedPerTpi
		}

		productionGraph, err := d.caughtRepository.GetProductionTotalGraphDashboard(tpiID, districtID, queryType, date)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetProductionGraphDashboard] Caught repository error")
		}

		productionValueGraph, err := d.auctionRepository.GetTransactionTotalGraphDashboard(tpiID, districtID, queryType, date)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetProductionValueGraphDashboard] Auction repository error")
		}

		transactionSpeedGraph, err := d.auctionRepository.GetTransactionSpeedGraphDashboard(tpiID, districtID, queryType, date)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetTransactionSpeedDashboard] Auction repository error")
		}

		result := map[string]interface{}{
			"production_total":        productionTotal,
			"transaction_total":       transactionTotal,
			"transaction_speed":       fmt.Sprintf("%.2f", transactionSpeed/3600),
			"production_total_graph":  productionGraph,
			"transaction_total_graph": productionValueGraph,
			"transaction_speed_graph": transactionSpeedGraph,
		}

		return result, nil
	}

	productionTotal, err := d.caughtRepository.GetProductionTotalDashboard(tpiID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetProductionTotalDashboard] Caught repository error")
	}

	transactionTotal, err := d.auctionRepository.GetTransactionValueDashboard(tpiID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetTransactionTotalDashboard] Auction repository error")
	}

	transactionSpeed, err := d.auctionRepository.GetTransactionSpeedDashboard(tpiID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
	}

	productionGraph, err := d.caughtRepository.GetProductionTotalGraphDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetProductionGraphDashboard] Caught repository error")
	}

	productionValueGraph, err := d.auctionRepository.GetTransactionTotalGraphDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetProductionValueGraphDashboard] Auction repository error")
	}

	transactionSpeedGraph, err := d.auctionRepository.GetTransactionSpeedGraphDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetTransactionSpeedDashboard] Auction repository error")
	}

	result := map[string]interface{}{
		"production_total":        productionTotal,
		"transaction_total":       transactionTotal,
		"transaction_speed":       fmt.Sprintf("%.2f", transactionSpeed/3600),
		"production_total_graph":  productionGraph,
		"transaction_total_graph": productionValueGraph,
		"transaction_speed_graph": transactionSpeedGraph,
	}

	return result, nil
}

func NewDashboardUsecase(caughRepository mysql.CaughtRepository, auctionRepository mysql.AuctionRepository, transactionRepository mysql.TransactionRepository, tpiRepository mysql.TpiRepository) DashboardUsecase {
	return &dashboardUsecase{caughtRepository: caughRepository, auctionRepository: auctionRepository, transactionRepository: transactionRepository, tpiRepository: tpiRepository}
}
