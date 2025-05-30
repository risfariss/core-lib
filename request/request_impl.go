package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/risfariss/core-lib/constant"
)

type Request struct {
	/*todo fill this block if need connection with another interface*/
}

func InitializeRequest() IRequest {
	return &Request{}
}

func (r *Request) CheckToken(ctx *gin.Context) (userId string, err error) {
	userId = ctx.MustGet("userId").(string)
	if userId != "" {
		return userId, nil
	}
	return "", errors.New(constant.ERROR_NOT_SEND_USER_ID)
}
