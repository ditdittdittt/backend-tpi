package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type UserHandler interface {
	Login(c *gin.Context)
	GetUser(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)
	Logout(c *gin.Context)
}

type userHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(server *gin.Engine, userUsecase usecase.UserUsecase) {
	handler := &userHandler{UserUsecase: userUsecase}
	user := server.Group("/auth")
	{
		user.POST("/get-user", middleware.AuthorizeJWT(constant.GetUser), handler.GetUser)
		user.POST("/login", handler.Login)
		user.POST("/logout", middleware.AuthorizeJWT(constant.Pass), handler.Logout)
	}
	server.GET("/users", handler.Index)
	server.GET("/user/:id", middleware.AuthorizeJWT(constant.GetByIDUser), handler.GetByID)
	server.PUT("/user/:id", middleware.AuthorizeJWT(constant.UpdateUser), handler.Update)
}

func (handler *userHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	token, err := handler.UserUsecase.Login(request.Username, request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
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
	curUserID := c.MustGet("userID")
	curUserIDint := curUserID.(int)

	user, location, err := handler.UserUsecase.GetUser(curUserIDint)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: map[string]interface{}{
			"user":     user,
			"location": location,
		},
	})
}

func (handler *userHandler) Index(c *gin.Context) {
	users, err := handler.UserUsecase.Index()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: users,
	})
}

func (handler *userHandler) Update(c *gin.Context) {

}

func (handler *userHandler) GetByID(c *gin.Context) {
	userID := c.Param("id")
	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	user, err := handler.UserUsecase.GetByID(intUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: user,
	})
}

func (handler *userHandler) Logout(c *gin.Context) {
	curUserID := c.MustGet("userID")
	curUserIDint := curUserID.(int)

	err := handler.UserUsecase.Logout(curUserIDint)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
