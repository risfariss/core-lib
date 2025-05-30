package core_lib

import (
	"github.com/risfariss/core-lib/request"
	"github.com/risfariss/core-lib/response"
)

type CoreLib struct {
	request  request.IRequest
	response response.IResponse
}

func InitializeCoreLib() *CoreLib {
	request := request.InitializeRequest()
	response := response.InitializeResponse()
	return &CoreLib{
		request:  request,
		response: response,
	}
}
