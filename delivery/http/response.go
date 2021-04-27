package http

import (
	"github.com/ditdittdittt/backend-tpi/constant"
)

type Response struct {
	ResponseCode string      `json:"response_code"`
	ResponseDesc string      `json:"response_desc"`
	ResponseData interface{} `json:"response_data"`
}

func NewResponse(code string, desc string, data interface{}) *Response {
	return &Response{
		ResponseCode: code,
		ResponseDesc: desc,
		ResponseData: data,
	}
}

func NewErrorResponse(err error) *Response {
	return &Response{
		ResponseCode: constant.ErrorResponseCode,
		ResponseDesc: constant.Failed,
		ResponseData: err.Error(),
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		ResponseCode: constant.SuccessResponseCode,
		ResponseDesc: constant.Success,
		ResponseData: data,
	}
}
