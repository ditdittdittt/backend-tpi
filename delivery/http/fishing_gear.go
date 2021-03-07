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

type FishingGearHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
}

type fishingGearHandler struct {
	FishingGearUsecase usecase.FishingGearUsecase
}

func NewFishingGearHandler(server *gin.Engine, fishingGearusecase usecase.FishingGearUsecase) {
	handler := &fishingGearHandler{FishingGearUsecase: fishingGearusecase}
	server.POST("/fishing-gear", middleware.AuthorizeJWT(constant.CreateFishingGear), handler.Create)
	server.GET("/fishing-gears", handler.Index)
	server.GET("/fishing-gear/:id", middleware.AuthorizeJWT(constant.GetByIDFishingGear), handler.GetByID)
	server.PUT("/fishing-gear/:id", middleware.AuthorizeJWT(constant.UpdateFishingGear), handler.Update)
	server.DELETE("/fishing-gear/:id", middleware.AuthorizeJWT(constant.DeleteFishingGear), handler.Delete)
}

func (h *fishingGearHandler) Create(c *gin.Context) {
	var request CreateFishingGearRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	fishingGear := &entities.FishingGear{
		Name: request.Name,
		Code: request.Code,
	}

	err := h.FishingGearUsecase.Create(fishingGear)
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

func (h *fishingGearHandler) Index(c *gin.Context) {
	fishingGears, err := h.FishingGearUsecase.Index()
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
		ResponseData: fishingGears,
	})
}

func (h *fishingGearHandler) Update(c *gin.Context) {
	var request UpdateFishingGearRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishingGearID := c.Param("id")
	intFishingGearID, err := strconv.Atoi(fishingGearID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishingGear := &entities.FishingGear{
		ID:   intFishingGearID,
		Name: request.Name,
		Code: request.Code,
	}

	err = h.FishingGearUsecase.Update(fishingGear)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *fishingGearHandler) Delete(c *gin.Context) {
	fishingGearID := c.Param("id")
	intFishingGearID, err := strconv.Atoi(fishingGearID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.FishingGearUsecase.Delete(intFishingGearID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *fishingGearHandler) GetByID(c *gin.Context) {
	fishingGearID := c.Param("id")
	intFishingGearID, err := strconv.Atoi(fishingGearID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishingGear, err := h.FishingGearUsecase.GetByID(intFishingGearID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: fishingGear,
	})
}
