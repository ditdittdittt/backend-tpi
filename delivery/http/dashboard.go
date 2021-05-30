package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type DashboardHandler interface {
	GetDashboardHeader(c *gin.Context)
	GetDashboardDetail(c *gin.Context)
}

type dashboardHandler struct {
	dashboardUsecase usecase.DashboardUsecase
}

func NewDashboardHandler(server *gin.Engine, dashboardUsecase usecase.DashboardUsecase) {
	handler := &dashboardHandler{dashboardUsecase: dashboardUsecase}
	api := server.Group("/api/v1")
	{
		api.GET("/dashboard/header", middleware.AuthorizeJWT(constant.Pass), handler.GetDashboardHeader)
		api.GET("/dashboard/detail", middleware.AuthorizeJWT(constant.Pass), handler.GetDashboardDetail)
	}
}

func (h *dashboardHandler) GetDashboardHeader(c *gin.Context) {
	_, tpiID, districtID, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	result, err := h.dashboardUsecase.GetFisherAndBuyerTotal(tpiID, districtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: result,
	})

}

func (h *dashboardHandler) GetDashboardDetail(c *gin.Context) {
	intTpiID := 0
	intDistrictID := 0

	tpiID, ok := c.Get("tpiID")
	if ok {
		intTpiID = tpiID.(int)
	}

	districtID, ok := c.Get("districtID")
	if ok {
		intDistrictID = districtID.(int)
	}

	queryType := c.DefaultQuery("query", "daily")

	result, err := h.dashboardUsecase.GetDashboardDetail(intTpiID, intDistrictID, queryType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: result,
	})
}
