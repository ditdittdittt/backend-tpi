package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type ReportHandler interface {
	Production(c *gin.Context)
	Transaction(c *gin.Context)
	ExportExcelProduction(c *gin.Context)
	ExportExcelTransaction(c *gin.Context)
	ExportPdfProduction(c *gin.Context)
	ExportPdfTransaction(c *gin.Context)
}

type reportHandler struct {
	reportUsecase usecase.ReportUsecase
}

func NewReportHandler(server *gin.Engine, reportUsecase usecase.ReportUsecase) {
	handler := &reportHandler{reportUsecase: reportUsecase}
	api := server.Group("/api/v1")
	{
		api.GET("/report/production", middleware.AuthorizeJWT(constant.Pass), handler.Production)
		api.GET("/report/transaction", middleware.AuthorizeJWT(constant.Pass), handler.Transaction)
		api.GET("/report/production/excel", middleware.AuthorizeJWT(constant.Pass), handler.ExportExcelProduction)
		api.GET("/report/transaction/excel", middleware.AuthorizeJWT(constant.Pass), handler.ExportExcelTransaction)
		api.GET("/report/production/pdf", middleware.AuthorizeJWT(constant.Pass), handler.ExportPdfProduction)
		api.GET("/report/transaction/pdf", middleware.AuthorizeJWT(constant.Pass), handler.ExportPdfTransaction)
	}
}

func (h *reportHandler) Production(c *gin.Context) {
	tpiID, districtID := h.getTpiAndDistrict(c)
	if tpiID == 0 && districtID == 0 {
		return
	}

	stringFrom, stringTo, _ := h.getDate(c)

	productionReport, err := h.reportUsecase.GetProductionReport(tpiID, districtID, stringFrom, stringTo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: productionReport,
	})
}

func (h *reportHandler) Transaction(c *gin.Context) {
	tpiID, districtID := h.getTpiAndDistrict(c)
	if tpiID == 0 && districtID == 0 {
		return
	}

	stringFrom, stringTo, _ := h.getDate(c)

	transactionReport, err := h.reportUsecase.GetTransactionReport(tpiID, districtID, stringFrom, stringTo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: transactionReport,
	})
}

func (h *reportHandler) ExportExcelProduction(c *gin.Context) {
	tpiID, districtID := h.getTpiAndDistrict(c)

	if tpiID == 0 && districtID == 0 {
		return
	}

	stringFrom, stringTo, queryType := h.getDate(c)

	xlsx, err := h.reportUsecase.ExportExcelProductionReport(tpiID, districtID, stringFrom, stringTo, queryType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Laporan Produksi Ikan.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func (h *reportHandler) ExportExcelTransaction(c *gin.Context) {
	tpiID, districtID := h.getTpiAndDistrict(c)
	if tpiID == 0 && districtID == 0 {
		return
	}

	stringFrom, stringTo, queryType := h.getDate(c)

	xlsx, err := h.reportUsecase.ExportExcelTransactionReport(tpiID, districtID, stringFrom, stringTo, queryType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Laporan Transaksi Lelang.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func (h *reportHandler) ExportPdfProduction(c *gin.Context) {
	tpiID, districtID := h.getTpiAndDistrict(c)

	if tpiID == 0 && districtID == 0 {
		return
	}

	stringFrom, stringTo, queryType := h.getDate(c)

	pdfProduction, err := h.reportUsecase.ExportPdfProductionReport(tpiID, districtID, stringFrom, stringTo, queryType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Laporan Produksi Ikan.pdf")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(pdfProduction)
}

func (h *reportHandler) ExportPdfTransaction(c *gin.Context) {
	tpiID, districtID := h.getTpiAndDistrict(c)

	if tpiID == 0 && districtID == 0 {
		return
	}

	stringFrom, stringTo, queryType := h.getDate(c)

	pdfTransaction, err := h.reportUsecase.ExportPdfTransactionReport(tpiID, districtID, stringFrom, stringTo, queryType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Laporan Transaksi Lelang.pdf")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write(pdfTransaction)
}

func (h *reportHandler) getDate(c *gin.Context) (string, string, string) {
	var from time.Time
	var to time.Time
	var queryType string

	daily, ok := c.GetQuery("daily")
	if ok {
		from, _ = time.Parse("2006-01-02", daily)
		to, _ = time.Parse("2006-01-02", daily)
		to = to.Add(24 * time.Hour)
		queryType = constant.Daily
	}

	monthly, ok := c.GetQuery("monthly")
	if ok {
		from, _ = time.Parse("2006-01", monthly)
		to, _ = time.Parse("2006-01", monthly)
		to = to.AddDate(0, 1, 0)
		queryType = constant.Monthly
	}

	yearly, ok := c.GetQuery("yearly")
	if ok {
		from, _ = time.Parse("2006", yearly)
		to, _ = time.Parse("2006", yearly)
		to = to.AddDate(1, 0, 0)
		queryType = constant.Yearly
	}

	period, ok := c.GetQuery("period")
	if ok {
		fromPeriod := strings.Split(period, ":")[0]
		toPeriod := strings.Split(period, ":")[1]
		from, _ = time.Parse("2006-01-02", fromPeriod)
		to, _ = time.Parse("2006-01-02", toPeriod)
		to = to.Add(24 * time.Hour)
		queryType = constant.Period
	}

	stringFrom := from.Format("2006-01-02 15:04:05")
	stringTo := to.Format("2006-01-02 15:04:05")

	return stringFrom, stringTo, queryType
}

func (h *reportHandler) getTpiAndDistrict(c *gin.Context) (int, int) {
	tpiID := c.DefaultQuery("tpi_id", "0")
	intTpiID, _ := strconv.Atoi(tpiID)

	intDistrictID := 0
	if intTpiID == 0 {
		districtID, ok := c.Get("districtID")
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(stacktrace.NewError("Missing request")))
			return 0, 0
		}
		intDistrictID = districtID.(int)
	}

	return intTpiID, intDistrictID
}
