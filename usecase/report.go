package usecase

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/leekchan/accounting"
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
)

type ReportUsecase interface {
	GetProductionReport(tpiID int, districtID int, from string, to string) (map[string]interface{}, error)
	GetTransactionReport(tpiID int, districtID int, from string, to string) (map[string]interface{}, error)
	ExportExcelProductionReport(tpiID int, districtID int, from string, to string, queryType string) (*excelize.File, error)
	ExportExcelTransactionReport(tpiID int, districtID int, from string, to string, queryType string) (*excelize.File, error)
	ExportPdfProductionReport(tpiID int, districtID int, from string, to string, queryType string) ([]byte, error)
	ExportPdfTransactionReport(tpiID int, districtID int, from string, to string, queryType string) ([]byte, error)
}

var queryMapping = map[string]interface{}{
	constant.Daily:   "Tanggal",
	constant.Monthly: "Bulan",
	constant.Yearly:  "Tahun",
	constant.Period:  "Tanggal",
}

type reportUsecase struct {
	caughtRepository          mysql.CaughtRepository
	auctionRepository         mysql.AuctionRepository
	transactionRepository     mysql.TransactionRepository
	fishTypeRepository        mysql.FishTypeRepository
	transactionItemRepository mysql.TransactionItemRepository
	tpiRepository             mysql.TpiRepository
	districtRepository        mysql.DistrictRepository
}

func NewReportUsecase(
	caughtRepository mysql.CaughtRepository,
	auctionRepository mysql.AuctionRepository,
	transactionRepository mysql.TransactionRepository,
	fishTypeRepository mysql.FishTypeRepository,
	transactionItemRepository mysql.TransactionItemRepository,
	tpiRepository mysql.TpiRepository,
	districtRepository mysql.DistrictRepository) ReportUsecase {
	return &reportUsecase{
		caughtRepository:          caughtRepository,
		auctionRepository:         auctionRepository,
		transactionRepository:     transactionRepository,
		fishTypeRepository:        fishTypeRepository,
		transactionItemRepository: transactionItemRepository,
		tpiRepository:             tpiRepository,
		districtRepository:        districtRepository,
	}
}

func (r *reportUsecase) ExportPdfProductionReport(tpiID int, districtID int, from string, to string, queryType string) ([]byte, error) {
	header, err := r.productionReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[productionReport] Report usecase error")
	}

	header["query_type"] = queryMapping[queryType]

	switch queryType {
	case constant.Daily:
		header["date"] = from[:10]
	case constant.Monthly:
		header["date"] = from[:7]
	case constant.Yearly:
		header["date"] = from[:4]
	case constant.Period:
		header["date"] = from[:10] + " : " + to[:10]
	}

	if districtID != 0 {
		district, err := r.districtRepository.Get(map[string]interface{}{"id": districtID})
		if err != nil {
			return nil, err
		}
		header["district_name"] = district[0].Name
	}

	if tpiID != 0 {
		tpi, err := r.tpiRepository.GetByID(tpiID)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		header["tpi_name"] = tpi.Name
		header["district_name"] = tpi.District.Name
	}

	data, err := r.getProductionTable(tpiID, districtID, from, to)
	if err != nil {
		return nil, err
	}

	pdf, err := r.exportPdfProductionReport(header, data)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}

func (r *reportUsecase) ExportPdfTransactionReport(tpiID int, districtID int, from string, to string, queryType string) ([]byte, error) {
	header, err := r.transactionReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[transactionReport] Report usecase error")
	}

	header["query_type"] = queryMapping[queryType]

	switch queryType {
	case constant.Daily:
		header["date"] = from[:10]
	case constant.Monthly:
		header["date"] = from[:7]
	case constant.Yearly:
		header["date"] = from[:4]
	case constant.Period:
		header["date"] = from[:10] + " : " + to[:10]
	}

	if districtID != 0 {
		district, err := r.districtRepository.Get(map[string]interface{}{"id": districtID})
		if err != nil {
			return nil, err
		}
		header["district_name"] = district[0].Name
	}

	if tpiID != 0 {
		tpi, err := r.tpiRepository.GetByID(tpiID)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		header["tpi_name"] = tpi.Name
		header["district_name"] = tpi.District.Name
	}

	data, err := r.getTransactionTable(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[getTransactionTable] Report usecase error")
	}

	pdf, err := r.exportPdfTransactionReport(header, data)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}

func (r *reportUsecase) ExportExcelTransactionReport(tpiID int, districtID int, from string, to string, queryType string) (*excelize.File, error) {
	header, err := r.transactionReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[transactionReport] Report usecase error")
	}

	header["query_type"] = queryMapping[queryType]

	switch queryType {
	case constant.Daily:
		header["date"] = from[:10]
	case constant.Monthly:
		header["date"] = from[:7]
	case constant.Yearly:
		header["date"] = from[:4]
	case constant.Period:
		header["date"] = from[:10] + " : " + to[:10]
	}

	if districtID != 0 {
		district, err := r.districtRepository.Get(map[string]interface{}{"id": districtID})
		if err != nil {
			return nil, err
		}
		header["district_name"] = district[0].Name
	}

	if tpiID != 0 {
		tpi, err := r.tpiRepository.GetByID(tpiID)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		header["tpi_name"] = tpi.Name
		header["district_name"] = tpi.District.Name
	}

	data, err := r.getTransactionTable(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[getTransactionTable] Report usecase error")
	}

	xlsx, err := r.exportExcelTransactionReport(header, data)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[exportExcelTransactionReport] Report usecase error")
	}

	return xlsx, nil
}

func (r *reportUsecase) ExportExcelProductionReport(tpiID int, districtID int, from string, to string, queryType string) (*excelize.File, error) {
	header, err := r.productionReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[productionReport] Report usecase error")
	}

	header["query_type"] = queryMapping[queryType]

	switch queryType {
	case constant.Daily:
		header["date"] = from[:10]
	case constant.Monthly:
		header["date"] = from[:7]
	case constant.Yearly:
		header["date"] = from[:4]
	case constant.Period:
		header["date"] = from[:10] + " : " + to[:10]
	}

	if districtID != 0 {
		district, err := r.districtRepository.Get(map[string]interface{}{"id": districtID})
		if err != nil {
			return nil, err
		}
		header["district_name"] = district[0].Name
	}

	if tpiID != 0 {
		tpi, err := r.tpiRepository.GetByID(tpiID)
		if err != nil {
			return nil, stacktrace.Propagate(err, "[GetByID] Tpi repository error")
		}
		header["tpi_name"] = tpi.Name
		header["district_name"] = tpi.District.Name
	}

	data, err := r.getProductionTable(tpiID, districtID, from, to)
	if err != nil {
		return nil, err
	}

	xlsx, err := r.exportExcelProductionReport(header, data)
	if err != nil {
		return nil, err
	}

	return xlsx, nil
}

func (r *reportUsecase) GetTransactionReport(tpiID int, districtID int, from string, to string) (map[string]interface{}, error) {
	data, err := r.transactionReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[transactionReport] Report usecase error")
	}

	transactionTable, err := r.getTransactionTable(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[getTransactionTable] Report usecase error")
	}

	data["transaction_table"] = transactionTable

	return data, nil
}

func (r *reportUsecase) GetProductionReport(tpiID int, districtID int, from string, to string) (map[string]interface{}, error) {
	data, err := r.productionReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, err
	}

	productionTable, err := r.getProductionTable(tpiID, districtID, from, to)
	if err != nil {
		return nil, err
	}

	data["production_table"] = productionTable

	return data, nil
}

func (r *reportUsecase) exportPdfProductionReport(header map[string]interface{}, data []map[string]interface{}) ([]byte, error) {
	pdfg, err := services.GeneratePdf(header, data, constant.ProductionPdf)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GeneratePdf] Services error")
	}
	return pdfg, nil
}

func (r *reportUsecase) exportExcelProductionReport(header map[string]interface{}, data []map[string]interface{}) (*excelize.File, error) {
	xlsx := excelize.NewFile()
	sheet1Name := "Laporan Produksi Ikan"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Laporan Produksi Ikan")
	xlsx.MergeCell(sheet1Name, "A1", "F1")
	xlsx.SetCellStyle(sheet1Name, "A1", "A1", generateXlsxStyle(xlsx))
	xlsx.SetCellValue(sheet1Name, "A2", "Nama Kabupaten")
	xlsx.SetCellValue(sheet1Name, "B2", header["district_name"])
	xlsx.SetCellValue(sheet1Name, "A3", "Nama TPI")
	xlsx.SetCellValue(sheet1Name, "B3", header["tpi_name"])

	xlsx.SetCellValue(sheet1Name, "A4", header["query_type"])
	xlsx.SetCellValue(sheet1Name, "B4", header["date"])

	xlsx.SetCellValue(sheet1Name, "A6", "Total Produksi")
	xlsx.SetCellValue(sheet1Name, "A7", "Nilai Produksi")
	xlsx.SetCellValue(sheet1Name, "A8", "Kecepatan Penjualan")
	xlsx.SetCellValue(sheet1Name, "B6", header["production_total"])
	xlsx.SetCellValue(sheet1Name, "B7", header["production_value"])
	xlsx.SetCellValue(sheet1Name, "B8", header["transaction_speed"])

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
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+12), each["production_total"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+12), each["production_value"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+12), each["transaction_speed"])
	}

	return xlsx, nil
}

func (r *reportUsecase) productionReport(tpiID int, districtID int, from string, to string) (map[string]interface{}, error) {
	data := map[string]interface{}{}

	if districtID != 0 {
		tpis, err := r.tpiRepository.Get(map[string]interface{}{"district_id": districtID})
		if err != nil {
			return data, stacktrace.Propagate(err, "[Get] Tpi repository error")
		}

		productionTotal := 0.0
		productionValue := 0.0
		avgTransactionSpeed := 0.0

		for _, tpi := range tpis {
			weightTotalPerTpi, err := r.caughtRepository.GetWeightTotal(0, tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
			}

			productionTotal += weightTotalPerTpi

			productionValuePerTpi, err := r.auctionRepository.GetPriceTotal(0, tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
			}

			productionValue += productionValuePerTpi

			averageTransactionSpeedPerTpi, err := r.auctionRepository.GetTransactionSpeed(0, tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
			}

			avgTransactionSpeed += averageTransactionSpeedPerTpi / 3600
		}

		data["production_total"] = weightFormatter(productionTotal)
		data["production_value"] = currencyFormatter(productionValue)
		data["transaction_speed"] = fmt.Sprintf("%.2f Jam", avgTransactionSpeed/float64(len(tpis)))

		return data, nil
	}

	weightTotal, err := r.caughtRepository.GetWeightTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
	}

	data["production_total"] = weightFormatter(weightTotal)

	productionValue, err := r.auctionRepository.GetPriceTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
	}

	data["production_value"] = currencyFormatter(productionValue)

	averageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
	}

	data["transaction_speed"] = fmt.Sprintf("%.2f Jam", averageTransactionSpeed/3600)

	return data, nil
}

func (r *reportUsecase) getProductionTable(tpiID int, districtID int, from string, to string) ([]map[string]interface{}, error) {
	productionTable := make([]map[string]interface{}, 0)

	queryMap := []string{"id", "name", "code"}
	fishTypes, err := r.fishTypeRepository.GetWithSelectedField(queryMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetWithSelectedField] Fish type repository error")
	}

	if districtID != 0 {
		tpis, err := r.tpiRepository.Get(map[string]interface{}{"district_id": districtID})
		if err != nil {
			return nil, stacktrace.Propagate(err, "[Get] Tpi repository error")
		}

		for _, fishType := range fishTypes {
			productionTotal := 0.0
			productionValue := 0.0
			avgTransactionSpeed := 0.0

			for _, tpi := range tpis {
				fishWeightTotal, err := r.caughtRepository.GetWeightTotal(fishType.ID, tpi.ID, from, to)
				if err != nil {
					return nil, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
				}

				fishProductionValue, err := r.auctionRepository.GetPriceTotal(fishType.ID, tpi.ID, from, to)
				if err != nil {
					return nil, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
				}

				fishAverageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(fishType.ID, tpi.ID, from, to)
				if err != nil {
					return nil, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
				}

				productionTotal += fishWeightTotal
				productionValue += fishProductionValue
				avgTransactionSpeed += fishAverageTransactionSpeed / 3600
			}

			data := map[string]interface{}{
				"id":                fishType.ID,
				"code":              fishType.Code,
				"name":              fishType.Name,
				"production_total":  weightFormatter(productionTotal),
				"production_value":  currencyFormatter(productionValue),
				"transaction_speed": fmt.Sprintf("%.2f Jam", avgTransactionSpeed/float64(len(tpis))),
			}
			productionTable = append(productionTable, data)
		}

		return productionTable, nil
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
			"production_total":  weightFormatter(fishWeightTotal),
			"production_value":  currencyFormatter(fishProductionValue),
			"transaction_speed": fmt.Sprintf("%.2f Jam", fishAverageTransactionSpeed/3600),
		}
		productionTable = append(productionTable, data)
	}

	return productionTable, nil
}

func (r *reportUsecase) exportPdfTransactionReport(header map[string]interface{}, data []map[string]interface{}) ([]byte, error) {
	pdfg, err := services.GeneratePdf(header, data, constant.TransactionPdf)
	if err != nil {
		return nil, err
	}
	return pdfg, nil
}

func (r *reportUsecase) exportExcelTransactionReport(header map[string]interface{}, data []map[string]interface{}) (*excelize.File, error) {
	xlsx := excelize.NewFile()
	sheet1Name := "Laporan Transaksi Lelang"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Laporan Transaksi Lelang")
	xlsx.MergeCell(sheet1Name, "A1", "H1")
	xlsx.SetCellStyle(sheet1Name, "A1", "A1", generateXlsxStyle(xlsx))
	xlsx.SetCellValue(sheet1Name, "A2", "Nama Kabupaten")
	xlsx.SetCellValue(sheet1Name, "B2", header["district_name"])
	xlsx.SetCellValue(sheet1Name, "A3", "Nama TPI")
	xlsx.SetCellValue(sheet1Name, "B3", header["tpi_name"])

	xlsx.SetCellValue(sheet1Name, "A4", header["query_type"])
	xlsx.SetCellValue(sheet1Name, "B4", header["date"])

	xlsx.SetCellValue(sheet1Name, "A6", "Jumlah Transaksi")
	xlsx.SetCellValue(sheet1Name, "A7", "Total Produksi")
	xlsx.SetCellValue(sheet1Name, "A8", "Nilai Produksi")
	xlsx.SetCellValue(sheet1Name, "A9", "Kecepatan Penjualan")
	xlsx.SetCellValue(sheet1Name, "B6", header["transaction_total"])
	xlsx.SetCellValue(sheet1Name, "B7", header["production_total"])
	xlsx.SetCellValue(sheet1Name, "B8", header["production_value"])
	xlsx.SetCellValue(sheet1Name, "B9", header["transaction_speed"])

	xlsx.SetCellValue(sheet1Name, "A10", "Jumlah nelayan yang mendaratkan ikan")
	xlsx.MergeCell(sheet1Name, "A10", "D10")
	xlsx.SetCellValue(sheet1Name, "A11", "Nelayan Tetap")
	xlsx.SetCellValue(sheet1Name, "B11", header["permanent_fisher"])
	xlsx.SetCellValue(sheet1Name, "A12", "Nelayan Pendatang")
	xlsx.SetCellValue(sheet1Name, "B12", header["temporary_fisher"])

	xlsx.SetCellValue(sheet1Name, "E10", "Jumlah pembeli yang ikut lelang ikan")
	xlsx.MergeCell(sheet1Name, "E10", "H10")
	xlsx.SetCellValue(sheet1Name, "E11", "Pembeli Tetap")
	xlsx.SetCellValue(sheet1Name, "F11", header["permanent_buyer"])
	xlsx.SetCellValue(sheet1Name, "E12", "Pembeli Pendatang")
	xlsx.SetCellValue(sheet1Name, "F12", header["temporary_buyer"])

	xlsx.SetCellValue(sheet1Name, "A14", "Tabel Transaksi Lelang")
	xlsx.MergeCell(sheet1Name, "A14", "H14")

	xlsx.SetCellValue(sheet1Name, "A15", "No")
	xlsx.SetCellValue(sheet1Name, "B15", "ID")
	xlsx.SetCellValue(sheet1Name, "C15", "Nama Nelayan")
	xlsx.SetCellValue(sheet1Name, "D15", "Nama Pembeli")
	xlsx.SetCellValue(sheet1Name, "E15", "Kode Ikan")
	xlsx.SetCellValue(sheet1Name, "F15", "Nama Ikan")
	xlsx.SetCellValue(sheet1Name, "G15", "Berat")
	xlsx.SetCellValue(sheet1Name, "H15", "Nilai Lelang (Rp)")

	err := xlsx.AutoFilter(sheet1Name, "A15", "H15", "")
	if err != nil {
		return xlsx, nil
	}

	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+16), i+1)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+16), each["code"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+16), each["fisher_name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+16), each["buyer_name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+16), each["fish_code"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+16), each["fish_name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+16), each["fix_weight"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", i+16), each["price"])
	}

	return xlsx, nil
}

func (r *reportUsecase) transactionReport(tpiID int, districtID int, from string, to string) (map[string]interface{}, error) {
	data := map[string]interface{}{}

	if districtID != 0 {
		tpis, err := r.tpiRepository.Get(map[string]interface{}{"district_id": districtID})
		if err != nil {
			return data, stacktrace.Propagate(err, "[Get] Tpi repository error")
		}

		transactionTotal := 0
		productionTotal := 0.0
		productionValue := 0.0
		avgTransactionSpeed := 0.0
		permanentFisher := 0
		temporaryFisher := 0
		permanentBuyer := 0
		temporaryBuyer := 0

		for _, tpi := range tpis {
			transactionTotalPerTpi, err := r.transactionRepository.GetTransactionTotal(tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetTotalTransaction] Total transaction error")
			}

			weightTotalPerTpi, err := r.caughtRepository.GetWeightTotal(0, tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
			}

			productionValuePerTpi, err := r.auctionRepository.GetPriceTotal(0, tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
			}

			averageTransactionSpeedPerTpi, err := r.auctionRepository.GetTransactionSpeed(0, tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
			}

			permanentFisherPerTpi, err := r.caughtRepository.GetFisherTotal("Tetap", tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetFisherTotal] Caught repository error")
			}

			temporaryFisherPerTpi, err := r.caughtRepository.GetFisherTotal("Pendatang", tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetFisherTotal] Caught repository error")
			}

			permanentBuyerPerTpi, err := r.transactionRepository.GetBuyerTotal("Tetap", tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetTotalBuyer] Transaction repository error")
			}

			temporaryBuyerPerTpi, err := r.transactionRepository.GetBuyerTotal("Pendatang", tpi.ID, from, to)
			if err != nil {
				return data, stacktrace.Propagate(err, "[GetTotalBuyer] Transaction repository error")
			}

			transactionTotal += transactionTotalPerTpi
			productionTotal += weightTotalPerTpi
			productionValue += productionValuePerTpi
			avgTransactionSpeed += averageTransactionSpeedPerTpi / 3600
			permanentFisher += permanentFisherPerTpi
			temporaryFisher += temporaryFisherPerTpi
			permanentBuyer += permanentBuyerPerTpi
			temporaryBuyer += temporaryBuyerPerTpi
		}

		data["transaction_total"] = totalFormatter(transactionTotal)
		data["production_total"] = weightFormatter(productionTotal)
		data["production_value"] = currencyFormatter(productionValue)
		data["transaction_speed"] = fmt.Sprintf("%.2f Jam", avgTransactionSpeed/float64(len(tpis)))
		data["permanent_fisher"] = peopleFormatter(permanentFisher)
		data["temporary_fisher"] = peopleFormatter(temporaryFisher)
		data["permanent_buyer"] = peopleFormatter(permanentBuyer)
		data["temporary_buyer"] = peopleFormatter(temporaryBuyer)

		return data, nil
	}

	transactionTotal, err := r.transactionRepository.GetTransactionTotal(tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTotalTransaction] Total transaction error")
	}

	data["transaction_total"] = totalFormatter(transactionTotal)

	weightTotal, err := r.caughtRepository.GetWeightTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetWeightTotal] Caught repository error")
	}

	data["production_total"] = weightFormatter(weightTotal)

	productionValue, err := r.auctionRepository.GetPriceTotal(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetPriceTotal] Auction repository error")
	}

	data["production_value"] = currencyFormatter(productionValue)

	averageTransactionSpeed, err := r.auctionRepository.GetTransactionSpeed(0, tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTransactionSpeed] Auction repository error")
	}

	data["transaction_speed"] = fmt.Sprintf("%.2f Jam", averageTransactionSpeed/3600)

	permanentFisher, err := r.caughtRepository.GetFisherTotal("Tetap", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetFisherTotal] Caught repository error")
	}

	data["permanent_fisher"] = peopleFormatter(permanentFisher)

	temporaryFisher, err := r.caughtRepository.GetFisherTotal("Pendatang", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetFisherTotal] Caught repository error")
	}

	data["temporary_fisher"] = peopleFormatter(temporaryFisher)

	permanentBuyer, err := r.transactionRepository.GetBuyerTotal("Tetap", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTotalBuyer] Transaction repository error")
	}

	data["permanent_buyer"] = peopleFormatter(permanentBuyer)

	temporaryBuyer, err := r.transactionRepository.GetBuyerTotal("Pendatang", tpiID, from, to)
	if err != nil {
		return data, stacktrace.Propagate(err, "[GetTotalBuyer] Transaction repository error")
	}

	data["temporary_buyer"] = peopleFormatter(temporaryBuyer)

	return data, nil
}

func (r *reportUsecase) getTransactionTable(tpiID int, districtID int, from string, to string) ([]map[string]interface{}, error) {
	transactionData, err := r.transactionItemRepository.GetReport(tpiID, districtID, from, to)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Get] Transaction item repository error")
	}

	for index, data := range transactionData {
		data["index"] = index + 1
		data["price"] = currencyFormatter(data["price"].(float64))
		data["fix_weight"] = modifyFormatter(data["weight_unit"].(string), data["weight"].(float64))
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

func currencyFormatter(input float64) string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ",", Format: "%s %v"}
	return ac.FormatMoneyFloat64(input)
}

func weightFormatter(input float64) string {
	ac := accounting.Accounting{Symbol: "Kg", Precision: 2, Thousand: ".", Decimal: ",", Format: "%v %s"}
	return ac.FormatMoneyFloat64(input)
}

func totalFormatter(input int) string {
	ac := accounting.Accounting{Symbol: "Transaksi", Thousand: ".", Format: "%v %s"}
	return ac.FormatMoneyInt(input)
}

func peopleFormatter(input int) string {
	ac := accounting.Accounting{Symbol: "Orang", Thousand: ".", Format: "%v %s"}
	return ac.FormatMoneyInt(input)
}

func modifyFormatter(symbol string, input float64) string {
	ac := accounting.Accounting{Symbol: symbol, Precision: 2, Thousand: ".", Decimal: ",", Format: "%v %s"}
	return ac.FormatMoneyFloat64(input)
}
