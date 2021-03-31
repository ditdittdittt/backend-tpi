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

type CaughtHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	Inquiry(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type caughtHandler struct {
	CaughtUsecase usecase.CaughtUsecase
}

func NewCaughtHandler(server *gin.Engine, caughtUsecase usecase.CaughtUsecase) {
	handler := &caughtHandler{CaughtUsecase: caughtUsecase}
	server.POST("/caught", middleware.AuthorizeJWT(constant.CreateCaught), handler.Create)
	server.GET("/caught/getbyid/:id", middleware.AuthorizeJWT(constant.GetByIDCaught), handler.GetByID)
	server.PUT("/caught/update/:id", middleware.AuthorizeJWT(constant.UpdateCaught), handler.Update)
	server.DELETE("/caught/delete/:id", middleware.AuthorizeJWT(constant.DeleteCaught), handler.Delete)
	server.GET("/caughts", middleware.AuthorizeJWT(constant.Pass), handler.Index)
	server.GET("/caught/inquiry", middleware.AuthorizeJWT(constant.InquiryCaught), handler.Inquiry)
}

func (h *caughtHandler) Create(c *gin.Context) {
	var request CreateCaughtRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	curUserID, curTpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	caught := &entities.Caught{
		UserID:        curUserID,
		TpiID:         curTpiID,
		FisherID:      request.FisherID,
		FishingGearID: request.FishingGearID,
		FishingAreaID: request.FishingAreaID,
		TripDay:       request.TripDay,
	}

	for _, item := range request.CaughtItems {
		caughtItem := &entities.CaughtItem{
			FishTypeID:     item.FishTypeID,
			Weight:         item.Weight,
			WeightUnit:     item.WeightUnit,
			CaughtStatusID: 1,
		}
		caught.CaughtItem = append(caught.CaughtItem, caughtItem)
	}

	err = h.CaughtUsecase.Create(caught)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *caughtHandler) Index(c *gin.Context) {
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

	caughts, err := h.CaughtUsecase.Index(intFisherID, intFishTypeID, intCaughtStatusID, tpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: caughts,
	})
}

func (h *caughtHandler) Inquiry(c *gin.Context) {
	fisherID := c.DefaultQuery("fisher_id", "0")
	intFisherID, _ := strconv.Atoi(fisherID)

	fishTypeID := c.DefaultQuery("fish_type_id", "0")
	intFishTypeID, _ := strconv.Atoi(fishTypeID)

	_, curTpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	caughts, err := h.CaughtUsecase.Inquiry(intFisherID, intFishTypeID, curTpiID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: caughts,
	})
}

func (h *caughtHandler) GetByID(c *gin.Context) {
	caughtID := c.Param("id")
	intCaughtID, err := strconv.Atoi(caughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	caught, err := h.CaughtUsecase.GetByID(intCaughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: caught,
	})
}

func (h *caughtHandler) Update(c *gin.Context) {
	var request UpdateCaughtRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	caughtID := c.Param("id")
	intCaughtID, err := strconv.Atoi(caughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	userID, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	caught := &entities.CaughtItem{
		ID:       intCaughtID,
		CaughtID: request.CaughtID,
		Caught: &entities.Caught{
			UserID:        userID,
			TpiID:         tpiID,
			FisherID:      request.FisherID,
			FishingGearID: request.FishingGearID,
			FishingAreaID: request.FishingAreaID,
			TripDay:       request.TripDay,
		},
		FishTypeID: request.FishTypeID,
		Weight:     request.Weight,
		WeightUnit: request.WeightUnit,
	}

	err = h.CaughtUsecase.Update(caught)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *caughtHandler) Delete(c *gin.Context) {
	caughtID := c.Param("id")
	intCaughtID, err := strconv.Atoi(caughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.CaughtUsecase.Delete(intCaughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
