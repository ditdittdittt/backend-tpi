package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type FisherHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
}

type fisherHandler struct {
	FisherUsecase usecase.FisherUsecase
}

func NewFisherHandler(server *gin.Engine, fisherUsecase usecase.FisherUsecase) {
	handler := &fisherHandler{FisherUsecase: fisherUsecase}
	server.POST("/fisher", middleware.AuthorizeJWT(constant.CreateFisher), handler.Create)
	server.GET("/fishers", middleware.AuthorizeJWT(constant.Pass), handler.Index)
	server.PUT("/fisher/:id", middleware.AuthorizeJWT(constant.UpdateFisher), handler.Update)
	server.GET("/fisher/:id", middleware.AuthorizeJWT(constant.GetByIDFisher), handler.GetByID)
	server.DELETE("/fisher/:id", middleware.AuthorizeJWT(constant.DeleteFisher), handler.Delete)
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

	userID, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fisher := &entities.Fisher{
		UserID:      userID,
		Nik:         request.Nik,
		Name:        request.Name,
		NickName:    request.NickName,
		Address:     request.Address,
		ShipType:    request.ShipType,
		AbkTotal:    request.AbkTotal,
		PhoneNumber: request.PhoneNumber,
	}

	err = h.FisherUsecase.Create(fisher, tpiID, request.Status)
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

func (h *fisherHandler) Index(c *gin.Context) {
	_, tpiID, _, err := helper.GetCurrentUserID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fishers, err := h.FisherUsecase.Index(tpiID)
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
		ResponseData: fishers,
	})
}

func (h *fisherHandler) Update(c *gin.Context) {
	var request UpdateFisherRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	curUserID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(stacktrace.NewError("Login status invalid")))
		return
	}

	fisherID := c.Param("id")
	intFisherID, err := strconv.Atoi(fisherID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fisher := &entities.Fisher{
		ID:          intFisherID,
		UserID:      curUserID.(int),
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		ShipType:    request.ShipType,
		AbkTotal:    request.AbkTotal,
		PhoneNumber: request.PhoneNumber,
		//Status:      request.Status,
	}

	err = h.FisherUsecase.Update(fisher)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *fisherHandler) GetByID(c *gin.Context) {
	fisherID := c.Param("id")
	intFisherID, err := strconv.Atoi(fisherID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	fisher, err := h.FisherUsecase.GetByID(intFisherID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: fisher,
	})
}

func (h *fisherHandler) Delete(c *gin.Context) {
	fisherID := c.Param("id")
	intFisherID, err := strconv.Atoi(fisherID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.FisherUsecase.Delete(intFisherID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
