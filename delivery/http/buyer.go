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

type BuyerHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
}

type buyerHandler struct {
	BuyerUsecase usecase.BuyerUsecase
}

func NewBuyerHandler(server *gin.Engine, buyerUsecase usecase.BuyerUsecase) {
	handler := &buyerHandler{BuyerUsecase: buyerUsecase}
	api := server.Group("/api/v1")
	{
		api.POST("/buyer", middleware.AuthorizeJWT(constant.CreateBuyer), handler.Create)
		api.GET("/buyers", middleware.AuthorizeJWT(constant.ReadBuyer), handler.Index)
		api.PUT("/buyer/:id", middleware.AuthorizeJWT(constant.UpdateBuyer), handler.Update)
		api.GET("/buyer/:id", middleware.AuthorizeJWT(constant.ReadBuyer), handler.GetByID)
		api.DELETE("/buyer/:id", middleware.AuthorizeJWT(constant.DeleteBuyer), handler.Delete)
	}
}

func (h *buyerHandler) Create(c *gin.Context) {
	var request CreateBuyerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	userID, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyer := &entities.Buyer{
		UserID:      userID,
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	err = h.BuyerUsecase.Create(buyer, tpiID, request.Status)
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

func (h *buyerHandler) Index(c *gin.Context) {
	_, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyers, err := h.BuyerUsecase.Index(tpiID)
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
		ResponseData: buyers,
	})
}

func (h *buyerHandler) Update(c *gin.Context) {
	var request UpdateBuyerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	userID, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyerID := c.Param("id")
	intBuyerID, err := strconv.Atoi(buyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyer := &entities.Buyer{
		ID:          intBuyerID,
		UserID:      userID,
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		TpiID:       tpiID,
	}

	err = h.BuyerUsecase.Update(buyer, request.Status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *buyerHandler) GetByID(c *gin.Context) {
	_, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyerID := c.Param("id")
	intBuyerID, err := strconv.Atoi(buyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyer, err := h.BuyerUsecase.GetByID(intBuyerID, tpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: buyer,
	})
}

func (h *buyerHandler) Delete(c *gin.Context) {
	buyerID := c.Param("id")
	intBuyerID, err := strconv.Atoi(buyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.BuyerUsecase.Delete(intBuyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
