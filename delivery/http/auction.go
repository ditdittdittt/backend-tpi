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

type AuctionHandler interface {
	Create(c *gin.Context)
	Inquiry(c *gin.Context)
}

type auctionHandler struct {
	auctionUsecase usecase.AuctionUsecase
}

func NewAuctionHandler(server *gin.Engine, auctionUsecase usecase.AuctionUsecase) {
	handler := &auctionHandler{auctionUsecase: auctionUsecase}
	server.POST("/auction", middleware.AuthorizeJWT(constant.CreateAuction), handler.Create)
	server.GET("/auction/inquiry", middleware.AuthorizeJWT(constant.InquiryAuction), handler.Inquiry)
}

func (h *auctionHandler) Create(c *gin.Context) {
	var request CreateAuctionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	curUserID := c.MustGet("userID")
	curTpiID := c.MustGet("tpiID")

	auction := &entities.Auction{
		UserID:   curUserID.(int),
		TpiID:    curTpiID.(int),
		CaughtID: request.CaughtID,
		Price:    request.Price,
	}

	err := h.auctionUsecase.Create(auction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *auctionHandler) Inquiry(c *gin.Context) {
	fisherID := c.DefaultQuery("fisher_id", "0")
	intFisherID, _ := strconv.Atoi(fisherID)

	fishTypeID := c.DefaultQuery("fish_type_id", "0")
	intFishTypeID, _ := strconv.Atoi(fishTypeID)

	tpiID := c.MustGet("tpiID")

	auctions, err := h.auctionUsecase.Inquiry(intFisherID, intFishTypeID, tpiID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: auctions,
	})
}
