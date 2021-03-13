package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type DistrictHandler interface {
	Index(c *gin.Context)
}

type districtHandler struct {
	districtUsecase usecase.DistrictUsecase
}

func NewDistrictHandler(server *gin.Engine, districtUsecase usecase.DistrictUsecase) {
	handler := districtHandler{districtUsecase: districtUsecase}
	server.GET("/districts", handler.Index)
}

func (h *districtHandler) Index(c *gin.Context) {
	provinceID, ok := c.GetQuery("province_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "province_id is missing"})
	}
	intProvinceID, _ := strconv.Atoi(provinceID)

	districts, err := h.districtUsecase.Index(intProvinceID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: districts,
	})
}
