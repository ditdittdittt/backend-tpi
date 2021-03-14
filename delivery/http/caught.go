package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/middleware"
	"github.com/ditdittdittt/backend-tpi/usecase"
)

type CaughtHandler interface {
	Create(c *gin.Context)
	Index(c *gin.Context)
	Inquiry(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type caughtHandler struct {
	CaughtUsecase usecase.CaughtUsecase
}

func NewCaughtHandler(server *gin.Engine, caughtUsecase usecase.CaughtUsecase) {
	handler := &caughtHandler{CaughtUsecase: caughtUsecase}
	server.POST("/caught", middleware.AuthorizeJWT(constant.CreateCaught), handler.Create)
	server.GET("/caught/inquiry", middleware.AuthorizeJWT(constant.InquiryCaught), handler.Inquiry)
	server.GET("/caught/getbyid/:id", middleware.AuthorizeJWT(constant.GetByIDCaught), handler.GetByID)
	server.PUT("/caught/update/:id", middleware.AuthorizeJWT(constant.UpdateCaught), handler.Update)
	server.DELETE("/caught/delete/:id", middleware.AuthorizeJWT(constant.DeleteCaught), handler.Delete)
	server.GET("/caughts", middleware.AuthorizeJWT(constant.Pass), handler.Index)
}

func (h *caughtHandler) Create(c *gin.Context) {
	var request CreateCaughtRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}

func (h *caughtHandler) Index(c *gin.Context) {
	fisherID := c.DefaultQuery("fisher_id", "0")
	intFisherID, _ := strconv.Atoi(fisherID)

	fishTypeID := c.DefaultQuery("fish_type_id", "0")
	intFishTypeID, _ := strconv.Atoi(fishTypeID)

	caughtStatusID := c.DefaultQuery("caught_status_id", "0")
	intCaughtStatusID, _ := strconv.Atoi(caughtStatusID)

	tpiID := c.MustGet("tpiID")

	caughts, err := h.CaughtUsecase.Index(intFisherID, intFishTypeID, intCaughtStatusID, tpiID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: caughts,
	})
}

func (h *caughtHandler) Inquiry(c *gin.Context) {
	fisherID := c.DefaultQuery("fisher_id", "0")
	intFisherID, _ := strconv.Atoi(fisherID)

	fishTypeID := c.DefaultQuery("fish_type_id", "0")
	intFishTypeID, _ := strconv.Atoi(fishTypeID)

	tpiID := c.MustGet("tpiID")

	caughts, err := h.CaughtUsecase.Inquiry(intFisherID, intFishTypeID, tpiID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: caughts,
	})
}

func (h *caughtHandler) GetByID(c *gin.Context) {
	caughtID := c.Param("id")
	intCaughtID, err := strconv.Atoi(caughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	caught, err := h.CaughtUsecase.GetByID(intCaughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: caught,
	})
}

func (h *caughtHandler) Update(c *gin.Context) {
	// TODO Update
	c.JSON(http.StatusOK, gin.H{"Message": "Fungsi belum diimplementasi"})

}

func (h *caughtHandler) Delete(c *gin.Context) {
	caughtID := c.Param("id")
	intCaughtID, err := strconv.Atoi(caughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	err = h.CaughtUsecase.Delete(intCaughtID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
	})
}
