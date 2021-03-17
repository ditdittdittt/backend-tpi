package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type ReportHandler interface {
	Production(c *gin.Context)
	Transaction(c *gin.Context)
	ExportExcelProduction(c *gin.Context)
	ExportExcelTransaction(c *gin.Context)
}

type reportHandler struct {
	reportUsecase usecase.ReportUsecase
}

func NewReportHandler(server *gin.Engine, reportUsecase usecase.ReportUsecase) {
	handler := &reportHandler{reportUsecase: reportUsecase}
	server.GET("/report/production", handler.Production)
	server.GET("/report/transaction", handler.Transaction)
	server.GET("/report/production/excel", handler.ExportExcelProduction)
	server.GET("/report/transaction/excel", handler.ExportExcelTransaction)
}

func (h *reportHandler) Production(c *gin.Context) {
	tpiID := c.DefaultQuery("tpi_id", "0")
	intTpiID, _ := strconv.Atoi(tpiID)

	stringFrom, stringTo := h.getDate(c)

	productionReport, err := h.reportUsecase.GetProductionReport(intTpiID, stringFrom, stringTo)
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
	tpiID := c.DefaultQuery("tpi_id", "0")
	intTpiID, _ := strconv.Atoi(tpiID)

	stringFrom, stringTo := h.getDate(c)

	transactionReport, err := h.reportUsecase.GetTransactionReport(intTpiID, stringFrom, stringTo)
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
	tpiID := c.DefaultQuery("tpi_id", "0")
	intTpiID, _ := strconv.Atoi(tpiID)

	stringFrom, stringTo := h.getDate(c)

	xlsx, err := h.reportUsecase.ExportExcelProductionReport(intTpiID, stringFrom, stringTo)
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
	tpiID := c.DefaultQuery("tpi_id", "0")
	intTpiID, _ := strconv.Atoi(tpiID)

	stringFrom, stringTo := h.getDate(c)

	xlsx, err := h.reportUsecase.ExportExcelTransactionReport(intTpiID, stringFrom, stringTo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Laporan Transaksi Lelang.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func (h *reportHandler) getDate(c *gin.Context) (string, string) {
	var from time.Time
	var to time.Time

	daily, ok := c.GetQuery("daily")
	if ok {
		from, _ = time.Parse("2006-01-02", daily)
		to, _ = time.Parse("2006-01-02", daily)
		to = to.Add(24 * time.Hour)
	}

	monthly, ok := c.GetQuery("monthly")
	if ok {
		from, _ = time.Parse("2006-01", monthly)
		to, _ = time.Parse("2006-01", monthly)
		to = to.AddDate(0, 1, 0)
	}

	yearly, ok := c.GetQuery("yearly")
	if ok {
		from, _ = time.Parse("2006", yearly)
		to, _ = time.Parse("2006", yearly)
		to = to.AddDate(1, 0, 0)
	}

	period, ok := c.GetQuery("period")
	if ok {
		fromPeriod := strings.Split(period, ":")[0]
		toPeriod := strings.Split(period, ":")[1]
		from, _ = time.Parse("2006-01-02", fromPeriod)
		to, _ = time.Parse("2006-01-02", toPeriod)
		to = to.Add(24 * time.Hour)
	}

	stringFrom := from.Format("2006-01-02 15:04:05")
	stringTo := to.Format("2006-01-02 15:04:05")

	return stringFrom, stringTo
}
