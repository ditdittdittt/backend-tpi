package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type UserDistrictHandler interface {
	CreateDistrictAdmin(c *gin.Context)
}

type userDistrictHandler struct {
	UserDistrictUsecase usecase.UserDistrictUsecase
}

func NewUserDistrictHandler(server *gin.Engine, userDistrictUsecase usecase.UserDistrictUsecase) {
	handler := &userDistrictHandler{UserDistrictUsecase: userDistrictUsecase}
	user := server.Group("/auth")
	{
		user.POST("/create-district-admin", middleware.AuthorizeJWT(constant.CreateDistrictAdmin), handler.CreateDistrictAdmin)
	}
}

func (h *userDistrictHandler) CreateDistrictAdmin(c *gin.Context) {
	var request CreateDistrictAdminRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	districtUser := &entities.UserDistrict{
		User: entities.User{
			RoleID:   2,
			Nik:      request.Nik,
			Name:     request.Name,
			Address:  request.Address,
			Username: request.Username,
		},
		DistrictID: request.DistrictID,
	}

	err := h.UserDistrictUsecase.CreateDistrictAccount(districtUser)
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
