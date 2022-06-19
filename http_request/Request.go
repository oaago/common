package http_request

import (
	"github.com/kirinlabs/HttpRequest"
)

func NewRequest() *HttpRequest.Request {
	// 具体示例 https://github.com/kirinlabs/HttpRequest
	req := HttpRequest.NewRequest()
	req.Debug(true).SetTimeout(50)
	return req
}
