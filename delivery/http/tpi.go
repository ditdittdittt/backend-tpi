package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type TpiHandler interface {
	Create(c *gin.Context)
}

type tpiHandler struct {
	TpiUsecase usecase.TpiUsecase
}

func NewTpiHandler(server *gin.Engine, tpiUsecase usecase.TpiUsecase) {
	handler := &tpiHandler{TpiUsecase: tpiUsecase}
	server.POST("/tpi", middleware.AuthorizeJWT(constant.CreateTpi), handler.Create)
}

func (h *tpiHandler) Create(c *gin.Context) {
	var request CreateTpiRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			ResponseCode: constant.ErrorResponseCode,
			ResponseDesc: constant.Failed,
			ResponseData: err.Error(),
		})
		return
	}

	districtID := c.MustGet("districtID")

	tpi := &entities.Tpi{
		DistrictID: districtID.(int),
		Name:       request.Name,
		Code:       request.Code,
	}

	err := h.TpiUsecase.Create(tpi)
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
