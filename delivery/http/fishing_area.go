package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type FishingAreaHandler interface {
	Create(c *gin.Context)
}

type fishingAreaHandler struct {
	fishingAreaUsecase usecase.FishingAreaUsecase
}

func NewFishingAreaHandler(server *gin.Engine, fishingAreaUsecase usecase.FishingAreaUsecase) {
	handler := &fishingAreaHandler{fishingAreaUsecase: fishingAreaUsecase}
	server.POST("/fishing-area", middleware.AuthorizeJWT(constant.CreateFishingArea), handler.Create)
}

func (h *fishingAreaHandler) Create(c *gin.Context) {
	var request CreateFishingAreaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	fishingArea := &entities.FishingArea{
		DistrictID:          request.DistrictID,
		SouthLatitudeDegree: request.SouthLatitudeDegree,
		SouthLatitudeMinute: request.SouthLatitudeMinute,
		SouthLatitudeSecond: request.SouthLatitudeSecond,
		EastLongitudeDegree: request.EastLongitudeDegree,
		EastLongitudeMinute: request.EastLongitudeMinute,
		EastLongitudeSecond: request.EastLongitudeSecond,
		Name:                request.Name,
	}

	err := h.fishingAreaUsecase.Create(fishingArea)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
