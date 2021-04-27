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
			return nil, err
		}

		permanentFisherTotal := 0
		permanentBuyerTotal := 0
		temporaryFisherTotal := 0
		temporaryBuyerTotal := 0

		for _, tpi := range tpis {
			permanentFisherTotalPerTpi, err := d.caughtRepository.GetFisherTotalDashboard(tpi.ID, constant.PermanentStatus)
			if err != nil {
				return nil, err
			}

			temporaryFisherTotalPerTpi, err := d.caughtRepository.GetFisherTotalDashboard(tpi.ID, constant.TemporaryStatus)
			if err != nil {
				return nil, err
			}

			permanentBuyerTotalPerTpi, err := d.transactionRepository.GetBuyerTotalDashboard(tpi.ID, constant.PermanentStatus)
			if err != nil {
				return nil, err
			}

			temporaryBuyerTotalPerTpi, err := d.transactionRepository.GetBuyerTotalDashboard(tpi.ID, constant.TemporaryStatus)
			if err != nil {
				return nil, err
			}

			permanentFisherTotal += permanentFisherTotalPerTpi
			permanentBuyerTotal += permanentBuyerTotalPerTpi
			temporaryFisherTotal += temporaryFisherTotalPerTpi
			temporaryBuyerTotal += temporaryBuyerTotalPerTpi
		}

		result := map[string]interface{}{
			"fisher_total": permanentFisherTotal + temporaryFisherTotal,
			"buyer_total":  permanentBuyerTotal + temporaryBuyerTotal,
			"fisher_total_status": []map[string]interface{}{
				{
					"status": constant.PermanentStatus,
					"total":  permanentFisherTotal,
				},
				{
					"status": constant.TemporaryStatus,
					"total":  temporaryFisherTotal,
				},
			},
			"buyer_total_status": []map[string]interface{}{
				{
					"status": constant.PermanentStatus,
					"total":  permanentBuyerTotal,
				},
				{
					"status": constant.TemporaryStatus,
					"total":  temporaryBuyerTotal,
				},
			},
		}

		return result, nil
	}

	permanentFisherTotal, err := d.caughtRepository.GetFisherTotalDashboard(tpiID, constant.PermanentStatus)
	if err != nil {
		return nil, err
	}

	temporaryFisherTotal, err := d.caughtRepository.GetFisherTotalDashboard(tpiID, constant.TemporaryStatus)
	if err != nil {
		return nil, err
	}

	permanentBuyerTotal, err := d.transactionRepository.GetBuyerTotalDashboard(tpiID, constant.PermanentStatus)
	if err != nil {
		return nil, err
	}

	temporaryBuyerTotal, err := d.transactionRepository.GetBuyerTotalDashboard(tpiID, constant.TemporaryStatus)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"fisher_total": permanentFisherTotal + temporaryFisherTotal,
		"buyer_total":  permanentBuyerTotal + temporaryBuyerTotal,
		"fisher_total_status": []map[string]interface{}{
			{
				"status": constant.PermanentStatus,
				"total":  permanentFisherTotal,
			},
			{
				"status": constant.TemporaryStatus,
				"total":  temporaryFisherTotal,
			},
		},
		"buyer_total_status": []map[string]interface{}{
			{
				"status": constant.PermanentStatus,
				"total":  permanentBuyerTotal,
			},
			{
				"status": constant.TemporaryStatus,
				"total":  temporaryBuyerTotal,
			},
		},
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
				return nil, err
			}

			transactionTotalPerTpi, err := d.auctionRepository.GetTransactionValueDashboard(tpi.ID, queryType, date)
			if err != nil {
				return nil, err
			}

			transactionSpeedPerTpi, err := d.auctionRepository.GetTransactionSpeedDashboard(tpi.ID, queryType, date)
			if err != nil {
				return nil, err
			}

			productionTotal += productionTotalPerTpi
			transactionTotal += transactionTotalPerTpi
			transactionSpeed += transactionSpeedPerTpi
		}

		productionGraph, err := d.caughtRepository.GetProductionTotalGraphDashboard(tpiID, districtID, queryType, date)
		if err != nil {
			return nil, err
		}

		productionValueGraph, err := d.auctionRepository.GetTransactionTotalGraphDashboard(tpiID, districtID, queryType, date)
		if err != nil {
			return nil, err
		}

		transactionSpeedGraph, err := d.auctionRepository.GetTransactionSpeedGraphDashboard(tpiID, districtID, queryType, date)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	transactionTotal, err := d.auctionRepository.GetTransactionValueDashboard(tpiID, queryType, date)
	if err != nil {
		return nil, err
	}

	transactionSpeed, err := d.auctionRepository.GetTransactionSpeedDashboard(tpiID, queryType, date)
	if err != nil {
		return nil, err
	}

	productionGraph, err := d.caughtRepository.GetProductionTotalGraphDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, err
	}

	productionValueGraph, err := d.auctionRepository.GetTransactionTotalGraphDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, err
	}

	transactionSpeedGraph, err := d.auctionRepository.GetTransactionSpeedGraphDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, err
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
