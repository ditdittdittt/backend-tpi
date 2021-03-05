package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type BuyerHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
}

type buyerHandler struct {
	BuyerUsecase usecase.BuyerUsecase
}

func NewBuyerHandler(server *gin.Engine, buyerUsecase usecase.BuyerUsecase) {
	handler := &buyerHandler{BuyerUsecase: buyerUsecase}
	server.POST("/buyer", middleware.AuthorizeJWT(constant.CreateBuyer), handler.Create)
	server.GET("/buyers", handler.Index)
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

	curUserID := c.MustGet("userID")

	buyer := &entities.Buyer{
		UserID:      curUserID.(int),
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Status:      request.Status,
	}

	err := h.BuyerUsecase.Create(buyer)
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
	buyers, err := h.BuyerUsecase.Index()
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
