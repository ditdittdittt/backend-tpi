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

type TransactionHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type transactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(server *gin.Engine, transactionUsecase usecase.TransactionUsecase) {
	handler := &transactionHandler{transactionUsecase: transactionUsecase}
	server.POST("/transaction", middleware.AuthorizeJWT(constant.CreateTransaction), handler.Create)
	server.GET("/transactions", middleware.AuthorizeJWT(constant.Pass), handler.Index)
	server.GET("/transaction/getbyid/:id", middleware.AuthorizeJWT(constant.GetByIDTransaction), handler.GetByID)
	server.PUT("/transaction/update/:id", middleware.AuthorizeJWT(constant.UpdateTransaction), handler.Update)
	server.DELETE("/transaction/delete/:id", middleware.AuthorizeJWT(constant.DeleteTransaction), handler.Delete)
}

func (h *transactionHandler) Create(c *gin.Context) {
	var request CreateTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	userID, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	transaction := &entities.Transaction{
		UserID:           userID,
		TpiID:            tpiID,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
	}

	err = h.transactionUsecase.Create(transaction, request.AuctionsIDs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *transactionHandler) Index(c *gin.Context) {
	_, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	transactions, err := h.transactionUsecase.Index(tpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: transactions,
	})
}

func (h *transactionHandler) GetByID(c *gin.Context) {
	transactionID := c.Param("id")
	intTransactionID, err := strconv.Atoi(transactionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	transaction, err := h.transactionUsecase.GetByID(intTransactionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: transaction,
	})

}

func (h *transactionHandler) Update(c *gin.Context) {
	var request UpdateTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	transactionID := c.Param("id")
	intTransactionID, err := strconv.Atoi(transactionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	userID, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	transaction := &entities.Transaction{
		ID:               intTransactionID,
		UserID:           userID,
		TpiID:            tpiID,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
	}

	err = h.transactionUsecase.Update(transaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *transactionHandler) Delete(c *gin.Context) {
	transactionID := c.Param("id")
	intTransactionID, err := strconv.Atoi(transactionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.transactionUsecase.Delete(intTransactionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
