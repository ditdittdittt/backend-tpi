package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type UserTpiHandler interface {
	CreateTpiAdmin(c *gin.Context)
	CreateTpiOfficer(c *gin.Context)
	CreateTpiCashier(c *gin.Context)
}

type usertTpiHandler struct {
	UserTpiUsecase usecase.UserTpiUsecase
}

func NewUserTpiHandler(server *gin.Engine, userTpiUsecase usecase.UserTpiUsecase)  {
	handler := &usertTpiHandler{UserTpiUsecase: userTpiUsecase}
	user := server.Group("/auth")
	{
		user.POST("/create-tpi-admin", middleware.AuthorizeJWT(constant.CreateTpiAdmin), handler.CreateTpiAdmin)
		user.POST("/create-tpi-officer", middleware.AuthorizeJWT(constant.CreateTpiOfficer), handler.CreateTpiOfficer)
		user.POST("/create-tpi-cashier", middleware.AuthorizeJWT(constant.CreateTpiCashier), handler.CreateTpiCashier)
	}
}

func (h *usertTpiHandler) CreateTpiAdmin(c *gin.Context) {
	var request CreateTpiAdminRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tpiUser := &entities.UserTpi{
		User:       entities.User{
			RoleID:   request.RoleID,
			Nik:      request.Nik,
			Name:     request.Name,
			Address:  request.Address,
			Username: request.Username,
		},
		TpiID: request.TpiID,
	}

	err := h.UserTpiUsecase.CreateTpiAccount(tpiUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			ResponseCode:    constant.ErrorResponseCode,
			ResponseDesc:    constant.Failed,
			ResponseData:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *usertTpiHandler) CreateTpiOfficer(c *gin.Context) {
	var request CreateTpiOfficerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tpiUser := &entities.UserTpi{
		User:       entities.User{
			RoleID:   request.RoleID,
			Nik:      request.Nik,
			Name:     request.Name,
			Address:  request.Address,
			Username: request.Username,
		},
		TpiID: request.TpiID,
	}

	err := h.UserTpiUsecase.CreateTpiAccount(tpiUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			ResponseCode:    constant.ErrorResponseCode,
			ResponseDesc:    constant.Failed,
			ResponseData:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *usertTpiHandler) CreateTpiCashier(c *gin.Context) {
	var request CreateTpiCashierRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tpiUser := &entities.UserTpi{
		User:       entities.User{
			RoleID:   request.RoleID,
			Nik:      request.Nik,
			Name:     request.Name,
			Address:  request.Address,
			Username: request.Username,
		},
		TpiID: request.TpiID,
	}

	err := h.UserTpiUsecase.CreateTpiAccount(tpiUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			ResponseCode:    constant.ErrorResponseCode,
			ResponseDesc:    constant.Failed,
			ResponseData:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}