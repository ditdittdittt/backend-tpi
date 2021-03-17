package usecase

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type ReportUsecase interface {
	GetProductionReport(tpiID int, from string, to string) (map[string]interface{}, error)
	GetTransactionReport(tpiID int, from string, to string) (map[string]interface{}, error)
	ExportExcelProductionReport(tpiID int, from string, to string) (*excelize.File, error)
	ExportExcelTransactionReport(tpiID int, from string, to string) (*excelize.File, error)
}

type reportUsecase struct {
	caughtRepository          mysql.CaughtRepository
	auctionRepository         mysql.AuctionRepository
	transactionRepository     mysql.TransactionRepository
	fishTypeRepository        mysql.FishTypeRepository
	transactionItemRepository mysql.TransactionItemRepository
	tpiRepository             mysql.TpiRepository
}

func NewReportUsecase(
	caughtRepository mysql.CaughtRepository,
	auctionRepository mysql.AuctionRepository,
	transactionRepository mysql.TransactionRepository,
	fishTypeRepository mysql.FishTypeRepository,
	transactionItemRepository mysql.TransactionItemRepository,
	tpiRepository mysql.TpiRepository) ReportUsecase {
	return &reportUsecase{
		caughtRepository:          caughtRepository,
		auctionRepository:         auctionRepository,
		transactionRepository:     transactionRepository,
		fishTypeRepository:        fishTypeRepository,
		transactionItemRepository: transactionItemRepository,
		tpiRepository:             tpiRepository,
	}
}

func (r *reportUsecase) ExportExcelTransactionReport(tpiID int, from string, to string) (*excelize.File, error) {
	header, err := r.transactionReport(tpiID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[transactionReport] Report usecase error")
	}

	header["from_date"] = from
	header["to_date"] = to

	tpi, err := r.tpiRepository.GetByID(tpiID)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
	}

	header["tpi_name"] = tpi.Name

	data, err := r.getTransactionTable(tpiID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[getTransactionTable] Report usecase error")
	}

	xlsx, err := r.exportExcelTransactionReport(header, data)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[exportExcelTransactionReport] Report usecase error")
	}

	return xlsx, nil
}

func (r *reportUsecase) ExportExcelProductionReport(tpiID int, from string, to string) (*excelize.File, error) {
	header, err := r.productionReport(tpiID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[productionReport] Report usecase error")
	}

	header["from_date"] = from
	header["to_date"] = to

	tpi, err := r.tpiRepository.GetByID(tpiID)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
	}

	header["tpi_name"] = tpi.Name

	data, err := r.getProductionTable(tpiID, from, to)
	if err != nil {
		return nil, err
	}

	xlsx, err := r.exportExcelProductionReport(header, data)
	if err != nil {
		return nil, err
	}

	return xlsx, nil
}

func (r *reportUsecase) GetTransactionReport(tpiID int, from string, to string) (map[string]interface{}, error) {
	data, err := r.transactionReport(tpiID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[transactionReport] Report usecase error")
	}

	transactionTable, err := r.getTransactionTable(tpiID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[getTransactionTable] Report usecase error")
	}

	data["transaction_table"] = transactionTable

	return data, nil
}

func (r *reportUsecase) GetProductionReport(tpiID int, from string, to string) (map[string]interface{}, error) {
	data, err := r.productionReport(tpiID, from, to)
	if err != nil {
		return nil, err
	}

	productionTable, err := r.getProductionTable(tpiID, from, to)
	if err != nil {
		return nil, err
	}

	data["production_table"] = productionTable

	return data, nil
}

func (r *reportUsecase) exportExcelProductionReport(header map[string]interface{}, data []map[string]interface{}) (*excelize.File, error) {
	xlsx := excelize.NewFile()
	sheet1Name := "Laporan Produksi"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Laporan Produksi Ikan")
	xlsx.MergeCell(sheet1Name, "A1", "F1")
	xlsx.SetCellStyle(sheet1Name, "A1", "A1", generateXlsxStyle(xlsx))
	xlsx.SetCellValue(sheet1Name, "A2", "Nama TPI")
	xlsx.SetCellValue(sheet1Name, "B2", header["tpi_name"])

	xlsx.SetCellValue(sheet1Name, "A3", "Tanggal")
	xlsx.SetCellValue(sheet1Name, "B3", fmt.Sprintf("%s - %s", header["from_date"], header["to_date"]))

	xlsx.SetCellValue(sheet1Name, "A5", "Total Produksi")
	xlsx.SetCellValue(sheet1Name, "A6", "Nilai Produksi")
	xlsx.SetCellValue(sheet1Name, "A7", "Kecepatan Penjualan")
	xlsx.SetCellValue(sheet1Name, "B5", fmt.Sprintf("%.2f Kg", header["production_total"]))
	xlsx.SetCellValue(sheet1Name, "B6", fmt.Sprintf("Rp %2.f", header["production_value"]))
	xlsx.SetCellValue(sheet1Name, "B7", fmt.Sprintf("%s Jam", header["transaction_speed"]))

	xlsx.SetCellValue(sheet1Name, "A10", "Tabel Produksi Ikan")
	xlsx.MergeCell(sheet1Name, "A10", "F10")

	xlsx.SetCellValue(sheet1Name, "A11", "No")
	xlsx.SetCellValue(sheet1Name, "B11", "Kode Ikan")
	xlsx.SetCellValue(sheet1Name, "C11", "Nama Ikan")
	xlsx.SetCellValue(sheet1Name, "D11", "Jumlah Produksi (Kg)")
	xlsx.SetCellValue(sheet1Name, "E11", "Nilai Produksi (Rp)")
	xlsx.SetCellValue(sheet1Name, "F11", "Rata-rata Kecepatan Penjualan (Jam)")

	err := xlsx.AutoFilter(sheet1Name, "A11", "F11", "")
	if err != nil {
		return xlsx, nil
	}

	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+12), i+1)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+12), each["code"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+12), each["name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+12), fmt.Sprintf("%.2f Kg", each["production_total"]))
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+12), fmt.Sprintf("Rp %2.f", each["production_value"]))
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+12), fmt.Sprintf("%s Jam", each["transaction_speed"]))
	}

	return xlsx, nil
}

func (r *reportUsecase) productionReport(tpiID int, from string, to string) (map[string]interface{}, error) {
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

	return data, nil
}

func (r *reportUsecase) getProductionTable(tpiID int, from string, to string) ([]map[string]interface{}, error) {
	productionTable := make([]map[string]interface{}, 0)
	queryMap := []string{"id", "name", "code"}
	fishTypes, err := r.fishTypeRepository.GetWithSelectedField(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetWithSelectedField] Fish type repository error")
	}

	for _, fishType := range fishTypes {
		fishWeightTotal, err := r.caughtRepository.GetWeightTotal(fishType.ID, tpiID, from, to)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
		}

		fishProductionValue, err := r.auctionRepository.GetPriceTotal(fishType.ID, tpiID, from, to)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
		}

		fishAverageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(fishType.ID, tpiID, from, to)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
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

	return productionTable, nil
}

func (r *reportUsecase) exportExcelTransactionReport(header map[string]interface{}, data []map[string]interface{}) (*excelize.File, error) {
	xlsx := excelize.NewFile()
	sheet1Name := "Laporan Transaksi"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Laporan Transaksi Lelang")
	xlsx.MergeCell(sheet1Name, "A1", "H1")
	xlsx.SetCellStyle(sheet1Name, "A1", "A1", generateXlsxStyle(xlsx))
	xlsx.SetCellValue(sheet1Name, "A2", "Nama TPI")
	xlsx.SetCellValue(sheet1Name, "B2", header["tpi_name"])

	xlsx.SetCellValue(sheet1Name, "A3", "Tanggal")
	xlsx.SetCellValue(sheet1Name, "B3", fmt.Sprintf("%s - %s", header["from_date"], header["to_date"]))

	xlsx.SetCellValue(sheet1Name, "A5", "Jumlah Transaksi")
	xlsx.SetCellValue(sheet1Name, "A6", "Total Produksi")
	xlsx.SetCellValue(sheet1Name, "A7", "Nilai Produksi")
	xlsx.SetCellValue(sheet1Name, "A8", "Kecepatan Penjualan")
	xlsx.SetCellValue(sheet1Name, "B5", fmt.Sprintf("%d transaksi", header["transaction_total"]))
	xlsx.SetCellValue(sheet1Name, "B6", fmt.Sprintf("%.2f Kg", header["production_total"]))
	xlsx.SetCellValue(sheet1Name, "B7", fmt.Sprintf("Rp %2.f", header["production_value"]))
	xlsx.SetCellValue(sheet1Name, "B8", fmt.Sprintf("%s Jam", header["transaction_speed"]))

	xlsx.SetCellValue(sheet1Name, "A9", "Jumlah nelayan yang mendaratkan ikan")
	xlsx.MergeCell(sheet1Name, "A9", "D9")
	xlsx.SetCellValue(sheet1Name, "A10", "Nelayan Tetap")
	xlsx.SetCellValue(sheet1Name, "B10", fmt.Sprintf("%d orang", header["permanent_fisher"]))
	xlsx.SetCellValue(sheet1Name, "A11", "Nelayan Pendatang")
	xlsx.SetCellValue(sheet1Name, "B11", fmt.Sprintf("%d orang", header["temporary_fisher"]))

	xlsx.SetCellValue(sheet1Name, "E9", "Jumlah pembeli yang ikut lelang ikan")
	xlsx.MergeCell(sheet1Name, "E9", "H9")
	xlsx.SetCellValue(sheet1Name, "E10", "Pembeli Tetap")
	xlsx.SetCellValue(sheet1Name, "F10", fmt.Sprintf("%d orang", header["permanent_buyer"]))
	xlsx.SetCellValue(sheet1Name, "E11", "Pembeli Pendatang")
	xlsx.SetCellValue(sheet1Name, "F11", fmt.Sprintf("%d orang", header["temporary_buyer"]))

	xlsx.SetCellValue(sheet1Name, "A13", "Tabel Transaksi Lelang")
	xlsx.MergeCell(sheet1Name, "A13", "H13")

	xlsx.SetCellValue(sheet1Name, "A14", "No")
	xlsx.SetCellValue(sheet1Name, "B14", "ID")
	xlsx.SetCellValue(sheet1Name, "C14", "Nama Nelayan")
	xlsx.SetCellValue(sheet1Name, "D14", "Nama Pembeli")
	xlsx.SetCellValue(sheet1Name, "E14", "Kode Ikan")
	xlsx.SetCellValue(sheet1Name, "F14", "Nama Ikan")
	xlsx.SetCellValue(sheet1Name, "G14", "Berat")
	xlsx.SetCellValue(sheet1Name, "H14", "Nilai Lelang (Rp)")

	err := xlsx.AutoFilter(sheet1Name, "A14", "H14", "")
	if err != nil {
		return xlsx, nil
	}

	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+15), i+1)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+15), each["code"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+15), each["fisher_name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+15), each["buyer_name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+15), each["fish_code"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+15), each["fish_name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+15), fmt.Sprintf("%.2f %s", each["weight"], each["weight_unit"]))
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", i+15), fmt.Sprintf("Rp %.2f", each["price"]))
	}

	return xlsx, nil
}

func (r *reportUsecase) transactionReport(tpiID int, from string, to string) (map[string]interface{}, error) {
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

	return data, nil
}

func (r *reportUsecase) getTransactionTable(tpiID int, from string, to string) ([]map[string]interface{}, error) {
	transactionData, err := r.transactionItemRepository.GetReport(tpiID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Transaction item repository error")
	}

	return transactionData, nil
}

func generateXlsxStyle(xlsx *excelize.File) int {
	style, _ := xlsx.NewStyle(`{
    "font": {
        "bold": true,
        "size": 28
    }
}`)
	xlsx.SetRowHeight(xlsx.GetSheetName(1), 1, 36)
	return style
}
