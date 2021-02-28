package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type AuctionHandler interface {
	Create(c *gin.Context)
}

type auctionHandler struct {
	auctionUsecase usecase.AuctionUsecase
}

func NewAuctionHandler(server *gin.Engine, auctionUsecase usecase.AuctionUsecase) {
	handler := &auctionHandler{auctionUsecase: auctionUsecase}
	server.POST("/auction", middleware.AuthorizeJWT(constant.CreateAuction), handler.Create)
}

func (h *auctionHandler) Create(c *gin.Context) {
	var request CreateAuctionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	auction := &entities.Auction{
		CaughtID: request.CaughtID,
		Price:    request.Price,
	}

	err := h.auctionUsecase.Create(auction)
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
