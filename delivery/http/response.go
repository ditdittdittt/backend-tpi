package http

type Response struct {
	ResponseCode	string		`json:"response_code"`
	ResponseDesc	string		`json:"response_desc"`
	ResponseData	interface{}	`json:"response_data"`
}

func NewResponse(code string, desc string, data interface{}) *Response {
	return &Response{
		ResponseCode: code,
		ResponseDesc: desc,
		ResponseData: data,
	}
}