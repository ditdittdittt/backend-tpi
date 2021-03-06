package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type TpiHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type tpiHandler struct {
	TpiUsecase usecase.TpiUsecase
}

func NewTpiHandler(server *gin.Engine, tpiUsecase usecase.TpiUsecase) {
	handler := &tpiHandler{TpiUsecase: tpiUsecase}
	api := server.Group("/api/v1")
	{
		api.POST("/tpi", middleware.AuthorizeJWT(constant.CreateTpi), handler.Create)
		api.GET("/tpis", middleware.AuthorizeJWT(constant.ReadTpi), handler.Index)
		api.GET("/tpi/:id", middleware.AuthorizeJWT(constant.ReadTpi), handler.GetByID)
		api.PUT("/tpi/:id", middleware.AuthorizeJWT(constant.UpdateTpi), handler.Update)
		api.DELETE("/tpi/:id", middleware.AuthorizeJWT(constant.DeleteTpi), handler.Delete)
	}
}

func (h *tpiHandler) Create(c *gin.Context) {
	var request CreateTpiRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	userID, _, districtID, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	tpi := &entities.Tpi{
		DistrictID:  districtID,
		UserID:      userID,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Pic:         request.Pic,
	}

	err = h.TpiUsecase.Create(tpi)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *tpiHandler) Index(c *gin.Context) {
	_, _, districtID, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	tpis, err := h.TpiUsecase.Index(districtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: tpis,
	})
}

func (h *tpiHandler) GetByID(c *gin.Context) {
	tpiID := c.Param("id")
	intTpiID, err := strconv.Atoi(tpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	tpi, err := h.TpiUsecase.GetByID(intTpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: tpi,
	})
}

func (h *tpiHandler) Update(c *gin.Context) {
	var request UpdateTpiRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	tpiID := c.Param("id")
	intTpiID, err := strconv.Atoi(tpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	userID, _, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	tpi := &entities.Tpi{
		ID:          intTpiID,
		UserID:      userID,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Pic:         request.Pic,
	}

	err = h.TpiUsecase.Update(tpi)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *tpiHandler) Delete(c *gin.Context) {
	tpiID := c.Param("id")
	intTpiID, err := strconv.Atoi(tpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.TpiUsecase.Delete(intTpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
