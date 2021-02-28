package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type UserHandler interface {
	Login(c *gin.Context)
	GetUser(c *gin.Context)
}

type userHandler struct {
	UserUsecase		usecase.UserUsecase
}

func NewUserHandler(server *gin.Engine, userUsecase usecase.UserUsecase)  {
	handler := &userHandler{UserUsecase: userUsecase}
	user := server.Group("/auth")
	{
		user.POST("/get-user", middleware.AuthorizeJWT(constant.GetUser), handler.GetUser)
		user.POST("/login", handler.Login)
	}
}

func (handler *userHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := handler.UserUsecase.Login(request.Username, request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			ResponseCode:    constant.ErrorResponseCode,
			ResponseDesc:    "Login error",
			ResponseData:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: "Login success",
		ResponseData: map[string]interface{}{
			"token": token,
		},
	})
}

func (handler *userHandler) GetUser(c *gin.Context) {
	var request GetUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curUserID := c.MustGet("userID")
	curUserIDint := curUserID.(int)

	user, err := handler.UserUsecase.GetUser(curUserIDint)
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
		ResponseData: map[string]interface{}{
			"id": user.ID,
			"role_id": user.RoleID,
			"nik": user.Nik,
			"name": user.Name,
			"address": user.Address,
			"username": user.Username,
		},
	})
}