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

type FishTypeHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
}

type fishTypeHandler struct {
	fishTypeUsecase usecase.FishTypeUsecase
}

func NewFishTypeHandler(server *gin.Engine, fishTypeUsecase usecase.FishTypeUsecase) {
	handler := &fishTypeHandler{fishTypeUsecase: fishTypeUsecase}
	server.POST("/fish-type", middleware.AuthorizeJWT(constant.CreateFishType), handler.Create)
	server.GET("/fish-types", handler.Index)
	server.GET("/fish-type/:id", middleware.AuthorizeJWT(constant.GetByIDFishType), handler.GetByID)
	server.PUT("/fish-type/:id", middleware.AuthorizeJWT(constant.UpdateFishType), handler.Update)
	server.DELETE("/fish-type/:id", middleware.AuthorizeJWT(constant.DeleteFishType), handler.Delete)
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

func (h *fishTypeHandler) Index(c *gin.Context) {
	fishTypes, err := h.fishTypeUsecase.Index()
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
		ResponseData: fishTypes,
	})
}

func (h *fishTypeHandler) Update(c *gin.Context) {
	var request UpdateFishTypeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishTypeID := c.Param("id")
	intFishTypeID, err := strconv.Atoi(fishTypeID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishType := &entities.FishType{
		ID:   intFishTypeID,
		Name: request.Name,
		Code: request.Code,
	}

	err = h.fishTypeUsecase.Update(fishType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *fishTypeHandler) GetByID(c *gin.Context) {
	fishTypeID := c.Param("id")
	intFishTypeID, err := strconv.Atoi(fishTypeID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishType, err := h.fishTypeUsecase.GetByID(intFishTypeID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: fishType,
	})
}

func (h *fishTypeHandler) Delete(c *gin.Context) {
	fishType := c.Param("id")
	intFishTypeID, err := strconv.Atoi(fishType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.fishTypeUsecase.Delete(intFishTypeID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
