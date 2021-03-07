package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type BuyerHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
}

type buyerHandler struct {
	BuyerUsecase usecase.BuyerUsecase
}

func NewBuyerHandler(server *gin.Engine, buyerUsecase usecase.BuyerUsecase) {
	handler := &buyerHandler{BuyerUsecase: buyerUsecase}
	server.POST("/buyer", middleware.AuthorizeJWT(constant.CreateBuyer), handler.Create)
	server.GET("/buyers", handler.Index)
	server.PUT("/buyer/:id", middleware.AuthorizeJWT(constant.UpdateBuyer), handler.Update)
	server.GET("/buyer/:id", middleware.AuthorizeJWT(constant.GetByIDBuyer), handler.GetByID)
	server.DELETE("/buyer/:id", middleware.AuthorizeJWT(constant.DeleteBuyer), handler.Delete)
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

func (h *buyerHandler) Update(c *gin.Context) {
	var request UpdateBuyerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	curUserID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(stacktrace.NewError("Login status invalid")))
		return
	}

	buyerID := c.Param("id")
	intBuyerID, err := strconv.Atoi(buyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyer := &entities.Buyer{
		ID:          intBuyerID,
		UserID:      curUserID.(int),
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Status:      request.Status,
	}

	err = h.BuyerUsecase.Update(buyer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *buyerHandler) GetByID(c *gin.Context) {
	buyerID := c.Param("id")
	intBuyerID, err := strconv.Atoi(buyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	buyer, err := h.BuyerUsecase.GetByID(intBuyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: buyer,
	})
}

func (h *buyerHandler) Delete(c *gin.Context) {
	buyerID := c.Param("id")
	intBuyerID, err := strconv.Atoi(buyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.BuyerUsecase.Delete(intBuyerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
