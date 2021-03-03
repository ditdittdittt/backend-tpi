package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type CaughtHandler interface {
	Create(c *gin.Context)
}

type caughtHandler struct {
	CaughtUsecase usecase.CaughtUsecase
}

func NewCaughtHandler(server *gin.Engine, caughtUsecase usecase.CaughtUsecase) {
	handler := &caughtHandler{CaughtUsecase: caughtUsecase}
	server.POST("/caught", middleware.AuthorizeJWT(constant.CreateCaught), handler.Create)
}

func (h *caughtHandler) Create(c *gin.Context) {
	var request CreateCaughtRequest
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

	caught := &entities.Caught{
		UserID:        curUserID.(int),
		TpiID:         curTpiID.(int),
		FisherID:      request.FisherID,
		FishingGearID: request.FishingGearID,
		FishingAreaID: request.FishingAreaID,
		TripDay:       request.TripDay,
	}

	err := h.CaughtUsecase.Create(caught, request.CaughtFishData)
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
