package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type FishTypeHandler interface {
	Create(c *gin.Context)
}

type fishTypeHandler struct {
	fishTypeUsecase usecase.FishTypeUsecase
}

func NewFishTypeHandler(server *gin.Engine, fishTypeUsecase usecase.FishTypeUsecase) {
	handler := &fishTypeHandler{fishTypeUsecase: fishTypeUsecase}
	server.POST("/fish-type", middleware.AuthorizeJWT(constant.CreateFishType), handler.Create)
}

func (h *fishTypeHandler) Create(c *gin.Context) {
	var request CreateFishTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	fishType := &entities.FishType{
		Name: request.Name,
		Code: request.Code,
	}

	err := h.fishTypeUsecase.Create(fishType)
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
