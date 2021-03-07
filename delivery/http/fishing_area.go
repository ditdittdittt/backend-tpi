package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type FishingAreaHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
}

type fishingAreaHandler struct {
	fishingAreaUsecase usecase.FishingAreaUsecase
}

func NewFishingAreaHandler(server *gin.Engine, fishingAreaUsecase usecase.FishingAreaUsecase) {
	handler := &fishingAreaHandler{fishingAreaUsecase: fishingAreaUsecase}
	server.POST("/fishing-area", middleware.AuthorizeJWT(constant.CreateFishingArea), handler.Create)
	server.GET("/fishing-areas", handler.Index)
	server.GET("/fishing-area/:id", middleware.AuthorizeJWT(constant.GetByIDFishingArea), handler.GetByID)
	server.PUT("/fishing-area/:id", middleware.AuthorizeJWT(constant.UpdateFishingArea), handler.Update)
	server.DELETE("/fishing-area/:id", middleware.AuthorizeJWT(constant.DeleteFishingArea), handler.Delete)
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

func (h *fishingAreaHandler) Index(c *gin.Context) {
	fishingAreas, err := h.fishingAreaUsecase.Index()
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
		ResponseData: fishingAreas,
	})
}

func (h *fishingAreaHandler) Update(c *gin.Context) {
	var request UpdateFishingAreaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishingAreaID := c.Param("id")
	intFishingAreaID, err := strconv.Atoi(fishingAreaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishingArea := &entities.FishingArea{
		ID:                  intFishingAreaID,
		Name:                request.Name,
		DistrictID:          request.DistrictID,
		SouthLatitudeDegree: request.SouthLatitudeDegree,
		SouthLatitudeMinute: request.SouthLatitudeMinute,
		SouthLatitudeSecond: request.SouthLatitudeSecond,
		EastLongitudeDegree: request.EastLongitudeDegree,
		EastLongitudeMinute: request.EastLongitudeMinute,
		EastLongitudeSecond: request.EastLongitudeSecond,
	}

	err = h.fishingAreaUsecase.Update(fishingArea)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *fishingAreaHandler) Delete(c *gin.Context) {
	fishingAreaID := c.Param("id")
	intFishingAreaID, err := strconv.Atoi(fishingAreaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.fishingAreaUsecase.Delete(intFishingAreaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *fishingAreaHandler) GetByID(c *gin.Context) {
	fishingAreaID := c.Param("id")
	intFishingAreaID, err := strconv.Atoi(fishingAreaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishingArea, err := h.fishingAreaUsecase.GetByID(intFishingAreaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: fishingArea,
	})
}
