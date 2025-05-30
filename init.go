package core_lib

import (
	"github.com/risfariss/core-lib/request"
	"github.com/risfariss/core-lib/response"
)

type CoreLib struct {
	Request  request.IRequest
	Response response.IResponse
}

func InitializeCoreLib() *CoreLib {
	request := request.InitializeRequest()
	response := response.InitializeResponse()
	return &CoreLib{
		Request:  request,
		Response: response,
	}
}
