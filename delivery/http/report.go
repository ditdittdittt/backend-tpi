package http

import "github.com/gin-gonic/gin"

type ReportHandler interface {
	Production(c *gin.Context)
}

type reportHandler struct {
}

func NewReportHandler(server *gin.Engine) {
	handler := &reportHandler{}
	server.GET("/report/production", handler.Production)
}

func (h *reportHandler) Production(c *gin.Context) {

}
