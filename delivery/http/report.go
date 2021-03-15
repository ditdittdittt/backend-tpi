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
}

type reportHandler struct {
	reportUsecase usecase.ReportUsecase
}

func NewReportHandler(server *gin.Engine, reportUsecase usecase.ReportUsecase) {
	handler := &reportHandler{reportUsecase: reportUsecase}
	server.GET("/report/production", handler.Production)
}

func (h *reportHandler) Production(c *gin.Context) {
	var from time.Time
	var to time.Time

	tpiID := c.DefaultQuery("tpi_id", "0")
	intTpiID, _ := strconv.Atoi(tpiID)

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

	productionReport, err := h.reportUsecase.Get(intTpiID, stringFrom, stringTo)
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
