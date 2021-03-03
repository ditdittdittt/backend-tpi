package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type TransactionHandler interface {
	Create(c *gin.Context)
}

type transactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(server *gin.Engine, transactionUsecase usecase.TransactionUsecase) {
	handler := &transactionHandler{transactionUsecase: transactionUsecase}
	server.POST("/transaction", middleware.AuthorizeJWT(constant.CreateTransaction), handler.Create)
}

func (h *transactionHandler) Create(c *gin.Context) {
	var request CreateTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	curUserID := c.MustGet("userID")
	curTpiID := c.MustGet("tpiID")

	transaction := &entities.Transaction{
		UserID:           curUserID.(int),
		TpiID:            curTpiID.(int),
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
		TotalPrice:       request.TotalPrice,
	}

	err := h.transactionUsecase.Create(transaction, request.AuctionsIDs)
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
