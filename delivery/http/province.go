package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type ProvinceHandler interface {
	Index(c *gin.Context)
}

type provinceHandler struct {
	provinceUsecase usecase.ProvinceUsecase
}

func NewProvinceHandler(server *gin.Engine, provinceUsecase usecase.ProvinceUsecase) {
	handler := provinceHandler{provinceUsecase: provinceUsecase}
	server.GET("/provinces", middleware.AuthorizeJWT(constant.Pass), handler.Index)
}

func (h *provinceHandler) Index(c *gin.Context) {
	provinces, err := h.provinceUsecase.Index()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: provinces,
	})
}
