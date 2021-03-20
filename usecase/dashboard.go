package usecase

import (
	"fmt"
	"time"

	"github.com/palantir/stacktrace"

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
}

var queryTypeMap = map[string]bool{
	"daily":   true,
	"monthly": true,
	"yearly":  true,
}

func (d *dashboardUsecase) GetFisherAndBuyerTotal(tpiID int, districtID int) (map[string]interface{}, error) {
	fisherTotal, err := d.caughtRepository.GetFisherTotalDashboard(tpiID, districtID)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetFisherTotalDashboard] Caught repository error")
	}

	buyerTotal, err := d.transactionRepository.GetBuyerTotalDashboard(tpiID, districtID)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetBuyerTotalDashboard] Transaction repository error")
	}

	result := map[string]interface{}{
		"fisher_total": fisherTotal,
		"buyer_total":  buyerTotal,
	}

	return result, nil
}

func (d *dashboardUsecase) GetDashboardDetail(tpiID int, districtID int, queryType string) (map[string]interface{}, error) {
	if !queryTypeMap[queryType] {
		return nil, stacktrace.NewError("Error invalid query type")
	}

	date := time.Now().Format("2006-01-02")

	productionTotal, err := d.caughtRepository.GetProductionTotalDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetProductionTotalDashboard] Caught repository error")
	}

	transactionTotal, err := d.auctionRepository.GetTransactionTotalDashboard(tpiID, districtID, queryType, date)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetTransactionTotalDashboard] Auction repository error")
	}

	transactionSpeed, err := d.auctionRepository.GetTransactionSpeedDashboard(tpiID, districtID, queryType, date)
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

func NewDashboardUsecase(caughRepository mysql.CaughtRepository, auctionRepository mysql.AuctionRepository, transactionRepository mysql.TransactionRepository) DashboardUsecase {
	return &dashboardUsecase{caughtRepository: caughRepository, auctionRepository: auctionRepository, transactionRepository: transactionRepository}
}
