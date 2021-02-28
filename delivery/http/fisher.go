package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type FisherHandler interface {
	Create(c *gin.Context)
}

type fisherHandler struct {
	FisherUsecase usecase.FisherUsecase
}

func NewFisherHandler(server *gin.Engine, fisherUsecase usecase.FisherUsecase) {
	handler := &fisherHandler{FisherUsecase: fisherUsecase}
	server.POST("/fisher", middleware.AuthorizeJWT(constant.CreateFisher), handler.Create)
}

func (h *fisherHandler) Create(c *gin.Context) {
	var request CreateFisherRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	curUserID := c.MustGet("userID")

	fisher := &entities.Fisher{
		UserID:      curUserID.(int),
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		ShipType:    request.ShipType,
		AbkTotal:    request.AbkTotal,
		PhoneNumber: request.PhoneNumber,
		Status:      request.Status,
	}

	err := h.FisherUsecase.Create(fisher)
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
