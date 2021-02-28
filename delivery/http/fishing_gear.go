package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type FishingGearHandler interface {
	Create(c *gin.Context)
}

type fishingGearHandler struct {
	FishingGearUsecase usecase.FishingGearUsecase
}

func NewFishingGearHandler(server *gin.Engine, fishingGearusecase usecase.FishingGearUsecase) {
	handler := &fishingGearHandler{FishingGearUsecase: fishingGearusecase}
	server.POST("/fishing-gear", middleware.AuthorizeJWT(constant.CreateFishingGear), handler.Create)
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
