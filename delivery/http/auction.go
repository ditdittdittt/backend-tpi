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

type AuctionHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	Inquiry(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type auctionHandler struct {
	auctionUsecase usecase.AuctionUsecase
}

func NewAuctionHandler(server *gin.Engine, auctionUsecase usecase.AuctionUsecase) {
	handler := &auctionHandler{auctionUsecase: auctionUsecase}
	api := server.Group("/api/v1")
	{
		api.POST("/auction", middleware.AuthorizeJWT(constant.CreateAuction), handler.Create)
		api.GET("/auction/inquiry", middleware.AuthorizeJWT(constant.ReadAuction), handler.Inquiry)
		api.GET("/auctions", middleware.AuthorizeJWT(constant.ReadAuction), handler.Index)
		api.GET("/auction/getbyid/:id}", middleware.AuthorizeJWT(constant.ReadAuction), handler.GetByID)
		api.PUT("/auction/update/:id", middleware.AuthorizeJWT(constant.UpdateAuction), handler.Update)
		api.DELETE("/auction/delete/:id", middleware.AuthorizeJWT(constant.DeleteAuction), handler.Delete)
	}
}

func (h *auctionHandler) Create(c *gin.Context) {
	var request CreateAuctionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	curUserID, curTpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	auction := &entities.Auction{
		UserID:       curUserID,
		TpiID:        curTpiID,
		CaughtItemID: request.CaughtItemID,
		Price:        request.Price,
	}

	err = h.auctionUsecase.Create(auction)
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

	_, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	auctions, err := h.auctionUsecase.Inquiry(intFisherID, intFishTypeID, tpiID)
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

func (h *auctionHandler) Index(c *gin.Context) {
	fisherID := c.DefaultQuery("fisher_id", "0")
	intFisherID, _ := strconv.Atoi(fisherID)

	fishTypeID := c.DefaultQuery("fish_type_id", "0")
	intFishTypeID, _ := strconv.Atoi(fishTypeID)

	caughtStatusID := c.DefaultQuery("caught_status_id", "0")
	intCaughtStatusID, _ := strconv.Atoi(caughtStatusID)

	_, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	auctions, err := h.auctionUsecase.Index(intFisherID, intFishTypeID, intCaughtStatusID, tpiID)
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

func (h *auctionHandler) GetByID(c *gin.Context) {
	auctionID := c.Param("id")
	intAuctionID, err := strconv.Atoi(auctionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	auction, err := h.auctionUsecase.GetByID(intAuctionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: auction,
	})
}

func (h *auctionHandler) Update(c *gin.Context) {
	var request UpdateAuctionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	auctionID := c.Param("id")
	intAuctionID, err := strconv.Atoi(auctionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	auction := &entities.Auction{
		ID:    intAuctionID,
		Price: request.Price,
	}

	err = h.auctionUsecase.Update(auction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *auctionHandler) Delete(c *gin.Context) {
	auctionID := c.Param("id")
	intAuctionID, err := strconv.Atoi(auctionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.auctionUsecase.Delete(intAuctionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
